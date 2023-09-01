package model

import (
	"github.com/kamva/mgm/v3"
	"peerswap/commodity/core/dto"
)

type Fiat struct {
	mgm.DefaultModel `bson:",inline"`
	Code             string `json:"code"`
	Name             string `json:"name"`
	CountryCode      string `json:"country_code"`
	Symbol           string `json:"symbol"`
	Icon             string `json:"icon"`
}

func (f Fiat) ToDto() *dto.Fiat {
	return &dto.Fiat{
		Id:          f.ID.Hex(),
		Code:        f.Code,
		Name:        f.Name,
		CountryCode: f.CountryCode,
		Symbol:      f.Symbol,
		Icon:        f.Icon,
		CratedAt:    f.CreatedAt,
		UpdatedAt:   f.UpdatedAt,
	}
}

func NewFiatFromFiatAddInput(input *dto.FiatAddInput) *Fiat {
	return &Fiat{
		Code:        input.Code,
		Name:        input.Name,
		CountryCode: input.CountryCode,
		Symbol:      input.Symbol,
		Icon:        input.Icon,
	}
}
