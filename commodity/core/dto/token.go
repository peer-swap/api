package dto

import "time"

type TokenAddInput struct {
	Address         string `json:"address"`
	ChainId         int    `json:"chain_id"`
	Decimals        int    `json:"decimals"`
	Symbol          string `json:"symbol"`
	Name            string `json:"name"`
	CoinGeckoId     string `json:"coin_gecko_id"`
	CoinMarketCapId string `json:"coin_market_cap_id"`
	Gasless         bool   `json:"gasless"`
	Icon            string `json:"icon"`
	MinimumAmount   string `json:"minimum_amount"`
	Fee             string `json:"fee"`
}

type Token struct {
	Id              string    `json:"id"`
	Address         string    `json:"address"`
	ChainId         int       `json:"chain_id"`
	Decimals        int       `json:"decimals"`
	Symbol          string    `json:"symbol"`
	Name            string    `json:"name"`
	CoinGeckoId     string    `json:"coin_gecko_id"`
	CoinMarketCapId string    `json:"coin_market_cap_id"`
	Gasless         bool      `json:"gasless"`
	Icon            string    `json:"icon"`
	MinimumAmount   string    `json:"minimum_amount"`
	Fee             string    `json:"fee"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type TokenExistsInput struct {
	ChainId int
	Address string
}

type TokenListFilter struct {
	Query   string `json:"query"`
	ChainId int    `json:"chain_id"`
}
