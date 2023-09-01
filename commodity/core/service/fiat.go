package service

import (
	"peerswap/commodity/core/dto"
	"peerswap/reusable"
)

type FiatDbInterface interface {
	Exists(string) (bool, error)
	Create(*dto.FiatAddInput) (*dto.Fiat, error)
	FindMany(string) ([]*dto.Fiat, error)
}

type Fiat struct {
	db FiatDbInterface
}

func (f Fiat) Add(input *dto.FiatAddInput) (*dto.Fiat, error) {
	if fails, err := reusable.NewValidator(input).Validate(); fails {
		return nil, err
	}

	if exists, err := f.db.Exists(input.Code); err != nil {
		return nil, DbErr
	} else if exists {
		return nil, CommodityExistsErr
	}

	fiat, err := f.db.Create(input)
	if err != nil {
		return nil, DbErr
	}

	return fiat, nil
}

func (f Fiat) List(query string) ([]*dto.Fiat, error) {
	fiats, err := f.db.FindMany(query)
	if err != nil {
		return nil, DbErr
	}

	if fiats == nil {
		return []*dto.Fiat{}, nil
	}

	return fiats, nil
}
