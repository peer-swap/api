package mongo

import (
	"github.com/kamva/mgm/v3"
	"peerswap/order/core/dto"
	"peerswap/reusable"
)

type Ad struct {
	mgm.DefaultModel `bson:",inline"`
	Type             reusable.TransactionType `bson:"type"`
	Asset            string                   `bson:"asset"`
	Fiat             string                   `bson:"fiat"`
	Price            float64                  `bson:"price"`
	Supply           float64                  `bson:"supply"`
	OrderUpperLimit  float64                  `bson:"orderUpperLimit"`
	OrderLowerLimit  float64                  `bson:"orderLowerLimit"`
	ChainId          uint                     `bson:"chainId"`
	Balance          float64                  `bson:"balance"`
	Status           reusable.AdStatus        `bson:"status"`
}

func (a Ad) ToDtoAd() *dto.Ad {
	return &dto.Ad{
		Id:              a.ID.Hex(),
		Type:            a.Type,
		Asset:           a.Asset,
		Fiat:            a.Fiat,
		Price:           a.Price,
		Supply:          a.Supply,
		OrderUpperLimit: a.OrderUpperLimit,
		OrderLowerLimit: a.OrderLowerLimit,
		ChainId:         a.ChainId,
		Balance:         a.Balance,
		Status:          a.Status,
		CreatedAt:       a.CreatedAt,
		UpdatedAt:       a.UpdatedAt,
	}
}
