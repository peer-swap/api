package mongo

import (
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"peerswap/ad/core/dto"
	"peerswap/ad/core/service"
)

type ModelPaymentMethodField struct {
	Id          string `bson:"id"`
	Name        string `bson:"name"`
	ContentType string `bson:"contentType"`
	IsRequired  bool   `bson:"isRequired"`
	IsCopyable  bool   `bson:"isCopyable"`
	IsDisplay   bool   `bson:"isDisplay"`
	Value       string `bson:"value"`
}
type ModelPaymentMethod struct {
	MethodId string                    `json:"methodId"`
	Display  string                    `json:"display"`
	Method   string                    `bson:"method"`
	Fields   []ModelPaymentMethodField `bson:"fields,omitempty"`
}

type Merchant struct {
	Id      primitive.ObjectID `bson:"id"`
	Address string             `bson:"address"`
}

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (s Adapter) UpdateBalance(id string, amount float64) (*dto.Ad, error) {
	adModel, err := s.find(id)
	if err != nil {
		return nil, err
	}
	adModel.Amount = amount
	err = mgm.Coll(adModel).Update(adModel)

	return adModel.ToDto(), err
}

func (s Adapter) Create(input dto.StoreInputDto) (*dto.Ad, error) {
	adModel := NewModelFromStoreInputDto(input)
	err := mgm.Coll(adModel).Create(adModel)
	if err != nil {
		return nil, service.DbError
	}

	return adModel.ToDto(), nil
}

func (s Adapter) Find(id string) (*dto.Ad, error) {
	m, err := s.find(id)
	if err != nil {
		return nil, err
	}
	return m.ToDto(), nil
}

func (s Adapter) UpdateActive(id string, active bool) (*dto.Ad, error) {
	m, err := s.find(id)
	if err != nil {
		return nil, err
	}

	m.Active = active

	err = mgm.Coll(m).Update(m)
	return m.ToDto(), err
}

func (s Adapter) List(input dto.ServiceListInputDto) ([]*dto.Ad, error) {
	var models []Ad
	query := bson.M{}
	if input.Type != "" {
		query["type"] = input.Type
	}
	if input.Fiat != "" {
		query["fiat"] = input.Fiat
	}
	if input.Asset != "" {
		query["asset"] = input.Asset
	}
	if input.Amount != 0 {
		query["orderLowerLimit"] = bson.M{operator.Lte: input.Amount}
		query["orderUpperLimit"] = bson.M{operator.Gte: input.Amount}
		query["balance"] = bson.M{operator.Gte: input.Amount}
	}
	if input.ChainId != 0 {
		query["chainId"] = input.ChainId
	}
	err := mgm.Coll(&Ad{}).SimpleFind(&models, query)
	if err != nil {
		return nil, err
	}

	var ads []*dto.Ad
	for _, model := range models {
		ads = append(ads, model.ToDto())
	}

	return ads, nil
}

func (s Adapter) find(id string) (*Ad, error) {
	m := &Ad{}

	err := mgm.Coll(m).FindByID(id, m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
