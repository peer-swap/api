package service

import (
	"errors"
	"peerswap/commodity/core/dto"
	"peerswap/reusable"
)

type TokenDbInterface interface {
	Exists(dto.TokenExistsInput) (bool, error)
	Create(*dto.TokenAddInput) (*dto.Token, error)
	FindMany(dto.TokenListFilter) ([]*dto.Token, error)
}

var (
	CommodityExistsErr = errors.New("token Exists in db")
	DbErr              = errors.New("db err")
)

type Token struct {
	db TokenDbInterface
}

func (t Token) Add(input *dto.TokenAddInput) (*dto.Token, error) {
	if fails, err := reusable.NewValidator(input).Validate(); fails {
		return nil, err
	}

	if exists, err := t.db.Exists(dto.TokenExistsInput{
		Address: input.Address,
		ChainId: input.ChainId,
	}); err != nil {
		return nil, DbErr
	} else if exists {
		return nil, CommodityExistsErr
	}

	token, err := t.db.Create(input)
	if err != nil {
		return nil, DbErr
	}

	return token, nil
}

func (t Token) List(filter dto.TokenListFilter) ([]*dto.Token, error) {
	tokens, err := t.db.FindMany(filter)
	if err != nil {
		return nil, DbErr
	}

	if tokens == nil {
		return []*dto.Token{}, nil
	}

	return tokens, nil
}
