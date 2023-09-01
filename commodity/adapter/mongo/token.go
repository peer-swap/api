package mongo

import (
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"peerswap/commodity/adapter/mongo/model"
	"peerswap/commodity/core/dto"
)

type Token struct {
}

func (a Token) FindMany(filter dto.TokenListFilter) ([]*dto.Token, error) {
	var _filter = bson.M{}
	var documents []*model.Token

	if filter.ChainId != 0 {
		_filter["chain_id"] = filter.ChainId
	}
	if filter.Query != "" {
		var or = bson.A{}
		regex := bson.M{operator.Regex: filter.Query, "$options": "i"}
		or = append(or, bson.M{"symbol": regex})
		or = append(or, bson.M{"name": bson.M{operator.Regex: filter.Query + ".*", "$options": "i"}})
		or = append(or, bson.M{"address": regex})
		or = append(or, bson.M{"coin_gecko_id": regex})
		or = append(or, bson.M{"coin_market_cap_id": regex})

		_filter[operator.Or] = or
	}

	if err := mgm.Coll(&model.Token{}).SimpleFind(documents, _filter); err != nil {
		return nil, err
	}

	var tokens []*dto.Token
	for _, token := range documents {
		tokens = append(tokens, token.ToDto())
	}

	return tokens, nil
}

func (a Token) Exists(input dto.TokenExistsInput) (bool, error) {
	m := &model.Token{}

	if err := mgm.Coll(m).First(bson.M{
		"chain_id": input.ChainId,
		"address":  input.Address,
	}, m); err != nil {
		return true, err
	}

	return m.ID.Hex() != "", nil
}

func (a Token) Create(input *dto.TokenAddInput) (*dto.Token, error) {
	m := model.NewTokenFromTokenAddInput(input)

	if err := mgm.Coll(m).Create(m); err != nil {
		return nil, err
	}

	return m.ToDto(), nil
}
