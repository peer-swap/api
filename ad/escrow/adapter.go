package escrow

import (
	"peerswap/ad/core/dto"
)

type Adapter struct {
}

func (a Adapter) VerifyTransaction(input *dto.AdFundingAddTokenInput) (*dto.FundingVerificationResponse, error) {
	return &dto.FundingVerificationResponse{
		Amount: input.Amount,
	}, nil
}

func (a Adapter) VerifySpendApproval(input *dto.AdFundingAddErc20Input) (*dto.FundingVerificationResponse, error) {
	return &dto.FundingVerificationResponse{
		Amount: input.Amount,
	}, nil
}

func (a Adapter) DepositErc20(input *dto.AdFundingAddErc20Input) error {
	return nil
}

func NewAdapter() *Adapter {
	return &Adapter{}
}
