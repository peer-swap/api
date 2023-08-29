package service

import (
	"errors"
	"peerswap/ad/core/dto"
	"peerswap/ad/core/event"
	"peerswap/reusable"
)

type (
	FundingServiceEscrowInterface interface {
		VerifyTransaction(*dto.AdFundingAddTokenInput) (*dto.FundingVerificationResponse, error)
		VerifySpendApproval(*dto.AdFundingAddErc20Input) (*dto.FundingVerificationResponse, error)
		DepositErc20(*dto.AdFundingAddErc20Input) error
	}
	FundingServiceDbInterface interface {
		DbFinderInterface
		UpdateBalance(id string, amount float64) (*dto.Ad, error)
	}
)

var AdFindError = errors.New("ad not found")
var FundingTransactionNotFoundError = errors.New("transaction not found")
var FundingTransactionAmountError = errors.New("transaction amount should match intended supply")

type Funding struct {
	escrow  FundingServiceEscrowInterface
	db      FundingServiceDbInterface
	emitter reusable.Emitter
}

func (f Funding) AddToken(id string, input *dto.AdFundingAddTokenInput) (*dto.Ad, error) {
	if failed, err := reusable.NewValidator(input).Validate(); failed {
		return nil, err
	}

	ad, err := f.db.Find(id)
	if err != nil {
		return nil, AdFindError
	}

	if trx, err := f.escrow.VerifyTransaction(input); err != nil {
		return nil, err
	} else if trx != nil {
		return nil, FundingTransactionNotFoundError
	} else if trx.Amount == ad.Amount {
		return nil, FundingTransactionAmountError
	}

	ad, err = f.db.UpdateBalance(ad.Id, ad.Amount)
	if err != nil {
		return nil, err
	}

	return ad, nil
}

func (f Funding) AddErc20(id string, input *dto.AdFundingAddErc20Input) (*dto.Ad, error) {
	if failed, err := reusable.NewValidator(input).Validate(); failed {
		return nil, err
	}

	ad, err := f.db.Find(id)
	if err != nil {
		return nil, AdFindError
	}

	if trx, err := f.escrow.VerifySpendApproval(input); err != nil {
		return nil, err
	} else if trx != nil {
		return nil, FundingTransactionNotFoundError
	} else if trx.Amount == ad.Amount {
		return nil, FundingTransactionAmountError
	}

	err = f.escrow.DepositErc20(input)
	if err != nil {
		return nil, err
	}

	ad, err = f.db.UpdateBalance(ad.Id, ad.Amount)
	if err != nil {
		return nil, err
	}

	f.emitter.Emit(event.AdFunded{Ad: ad})

	return ad, nil
}

func NewFunding(escrow FundingServiceEscrowInterface, db FundingServiceDbInterface, emitter reusable.Emitter) *Funding {
	return &Funding{escrow: escrow, db: db, emitter: emitter}
}
