package order

import (
	"errors"
	"peerswap/order/dto"
	"peerswap/reusable"
)

type DbInterface interface {
	create(*dto.ServiceStoreInput) (*dto.Order, error)
	findAd(*dto.ServiceStoreInput) (*dto.Ad, error)
	decrementAdBalance(string, float64) (*dto.Ad, error)
}

type EscrowInterface interface {
	PlaceOrder(*dto.ServiceStoreInput) error
}

type Service struct {
	db     DbInterface
	escrow EscrowInterface
}

var NotFoundError = errors.New("not found")

func (s Service) store(input *dto.ServiceStoreInput) (*dto.Order, *dto.Ad, error) {
	if fails, err := reusable.NewValidator(input).Validate(); fails {
		return nil, nil, err
	}

	ad, err := s.db.findAd(input)
	if err != nil {
		return nil, nil, err
	}

	order, err := s.db.create(input)
	if err != nil {
		return nil, nil, err
	}

	err = s.escrow.PlaceOrder(input)
	if err != nil {
		return nil, nil, err
	}
	ad, err = s.db.decrementAdBalance(input.Ad, input.Amount)
	if err != nil {
		return nil, nil, err
	}

	return order, ad, nil
}

func NewService(db DbInterface, escrow EscrowInterface) *Service {
	return &Service{db: db, escrow: escrow}
}
