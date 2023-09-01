package mongo

import (
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"peerswap/order/core"
	"peerswap/order/core/dto"
	"peerswap/reusable"
	"time"
)

type Adapter struct {
}

func (m Adapter) Find(id string) (*dto.Order, error) {
	order, err := m.find(id)
	if err != nil {
		return nil, err
	}

	return order.ToDtoOrder(), nil
}

func (m Adapter) FindAd(input *dto.ServiceStoreInput) (*dto.Ad, error) {
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
		return nil, core.NotFoundError
	}

	return ad.ToDtoAd(), nil
}

func (m Adapter) DecrementAdBalance(adId string, f float64) (*dto.Ad, error) {
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

func (m Adapter) Create(input *dto.ServiceStoreInput) (*dto.Order, error) {
	o := NewOrderFromDtoStoreInput(input)
	o.Status = reusable.OrderStatusPending
	err := mgm.Coll(o).Create(o)
	if err != nil {
		return nil, err
	}

	return o.ToDtoOrder(), nil
}

func (m Adapter) UpdateStatus(id string, status reusable.OrderStatus) (*dto.Order, error) {
	if status == reusable.OrderStatusCanceled || status == reusable.OrderStatusAppealed {
		return nil, core.CantUpdateErr
	}
	o, err := m.find(id)
	if err != nil {
		return nil, err
	}

	o.Status = status
	switch status {
	case reusable.OrderStatusPaymentSent:
		o.NotifiedPaidAt = time.Now()
		break
	case reusable.OrderStatusPaymentReceived:
		o.ConfirmPaidAt = time.Now()
		break
	case reusable.OrderStatusCompleted:
		o.ConfirmPaidAt = time.Now()
		break
	}

	return o.ToDtoOrder(), nil
}

func (m Adapter) Cancel(id string, input *dto.ServiceCancelInput) (*dto.Order, error) {
	o, err := m.find(id)
	if err != nil {
		return nil, err
	}

	o.Status = reusable.OrderStatusCanceled
	o.Canceled = Canceled{
		At:     time.Now(),
		By:     input.By,
		Reason: input.Reason,
	}

	return o.ToDtoOrder(), nil
}

func (m Adapter) Appeal(id string, input *dto.ServiceAppealInput) (*dto.Order, error) {
	o, err := m.find(id)
	if err != nil {
		return nil, err
	}

	o.Status = reusable.OrderStatusAppealed
	o.Appealed = Appealed{
		At:          time.Now(),
		By:          input.By,
		Reason:      input.Reason,
		Attachments: input.Attachments,
	}

	return o.ToDtoOrder(), nil
}

func (m Adapter) find(id string) (*Order, error) {
	o := &Order{}

	err := mgm.Coll(o).FindByID(id, o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func NewAdapter() *Adapter {
	return &Adapter{}
}
