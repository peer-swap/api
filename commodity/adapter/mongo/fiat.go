package mongo

import (
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"peerswap/commodity/adapter/mongo/model"
	"peerswap/commodity/core/dto"
)

type Fiat struct {
}

func (f Fiat) FindMany(code string) ([]*dto.Fiat, error) {
	var _filter = bson.M{}
	var documents []*model.Fiat

	if code != "" {
		var or = bson.A{}
		regex := bson.M{operator.Regex: code, "$options": "i"}
		or = append(or, bson.M{"code": regex})
		or = append(or, bson.M{"country_code": regex})
		or = append(or, bson.M{"name": bson.M{operator.Regex: code + ".*", "$options": "i"}})

		_filter[operator.Or] = or
	}

	if err := mgm.Coll(&model.Fiat{}).SimpleFind(documents, _filter); err != nil {
		return nil, err
	}

	var fiats []*dto.Fiat
	for _, token := range documents {
		fiats = append(fiats, token.ToDto())
	}

	return fiats, nil
}

func (f Fiat) Exists(code string) (bool, error) {
	m := &model.Fiat{}

	if err := mgm.Coll(m).First(bson.M{"code": code}, m); err != nil {
		return true, err
	}

	return m.ID.Hex() != "", nil
}

func (f Fiat) Create(input *dto.FiatAddInput) (*dto.Fiat, error) {
	m := model.NewFiatFromFiatAddInput(input)

	if err := mgm.Coll(m).Create(m); err != nil {
		return nil, err
	}

	return m.ToDto(), nil
}
