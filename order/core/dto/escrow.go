package dto

import "peerswap/reusable"

type PlaceOrderInput struct {
	Ad     string                   `json:"ad" validate:"required,alphanum"`
	Amount float64                  `json:"amount" validate:"required,numeric"`
	Type   reusable.TransactionType `json:"type" validate:"required,alpha,oneof=BUY SELL"`
	Asset  string                   `json:"asset" validate:"required,alphanum"`
	Fiat   string                   `json:"fiat" validate:"required,alphanum"`
	Seller PlaceOrderInputUser      `json:"seller,omitempty"`
	Buyer  PlaceOrderInputUser      `json:"buyer,omitempty"`
}

type PlaceOrderInputUser struct {
	Address string `json:"address" validate:"required,alphanum"`
}
