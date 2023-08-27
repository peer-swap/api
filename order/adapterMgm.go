package order

import (
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"peerswap/order/dto"
	"peerswap/reusable"
)

type AdapterMgm struct {
}

func (m AdapterMgm) findAd(input *dto.ServiceStoreInput) (*dto.Ad, error) {
	var ad = &Ad{}
	objectID, err := primitive.ObjectIDFromHex(input.Ad)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"_id":             objectID,
		"fiat":            input.Fiat,
		"asset":           input.Asset,
		"type":            input.Type,
		"orderUpperLimit": bson.M{operator.Gte: input.Amount},
		"orderLowerLimit": bson.M{operator.Lte: input.Amount},
	}
	if err := mgm.Coll(ad).First(filter, ad); err != nil {
		return nil, err
	} else if ad == nil {
		return nil, NotFoundError
	}

	return ad.ToDtoAd(), nil
}

func (m AdapterMgm) decrementAdBalance(adId string, f float64) (*dto.Ad, error) {
	var ad = &Ad{}
	if err := mgm.Coll(ad).FindByID(adId, ad); err != nil {
		return nil, err
	}

	ad.Balance -= f
	if err := mgm.Coll(ad).Update(ad); err != nil {
		return nil, err
	}

	return ad.ToDtoAd(), nil
}

func (m AdapterMgm) create(input *dto.ServiceStoreInput) (*dto.Order, error) {
	order := NewOrderFromDtoStoreInput(input)
	order.Status = reusable.OrderStatusPending
	err := mgm.Coll(order).Create(order)
	if err != nil {
		return nil, err
	}

	return order.ToDtoOrder(), nil
}

func NewAdapterMgm() *AdapterMgm {
	return &AdapterMgm{}
}
