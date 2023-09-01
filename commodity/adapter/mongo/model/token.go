package model

import (
	"github.com/kamva/mgm/v3"
	"peerswap/commodity/core/dto"
)

type Token struct {
	mgm.DefaultModel `bson:",inline"`
	Address          string `bson:"address"`
	ChainId          int    `bson:"chain_id"`
	Decimals         int    `bson:"decimals"`
	Symbol           string `bson:"symbol"`
	Name             string `bson:"name"`
	CoinGeckoId      string `bson:"coin_gecko_id"`
	CoinMarketCapId  string `bson:"coin_market_cap_id"`
	Gasless          bool   `bson:"gasless"`
	Icon             string `bson:"icon"`
	MinimumAmount    string `bson:"minimum_amount"`
	Fee              string `bson:"fee"`
}

func NewTokenFromTokenAddInput(t *dto.TokenAddInput) *Token {
	return &Token{
		Address:         t.Address,
		ChainId:         t.ChainId,
		Decimals:        t.Decimals,
		Symbol:          t.Symbol,
		Name:            t.Name,
		CoinGeckoId:     t.CoinGeckoId,
		CoinMarketCapId: t.CoinMarketCapId,
		Gasless:         t.Gasless,
		Icon:            t.Icon,
		MinimumAmount:   t.MinimumAmount,
		Fee:             t.Fee,
	}
}

func (t Token) ToDto() *dto.Token {
	return &dto.Token{
		Id:              t.ID.Hex(),
		Address:         t.Address,
		ChainId:         t.ChainId,
		Decimals:        t.Decimals,
		Symbol:          t.Symbol,
		Name:            t.Name,
		CoinGeckoId:     t.CoinGeckoId,
		CoinMarketCapId: t.CoinMarketCapId,
		Gasless:         t.Gasless,
		Icon:            t.Icon,
		MinimumAmount:   t.MinimumAmount,
		Fee:             t.Fee,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
	}
}
