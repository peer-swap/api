package dto

import (
	"peerswap/reusable"
	"time"
)

type Ad struct {
	Id              string                   `json:"id"`
	Type            reusable.TransactionType `json:"type"`
	Asset           string                   `json:"asset"`
	Fiat            string                   `json:"fiat"`
	Price           float64                  `json:"price"`
	PaymentMethods  []PaymentMethod          `json:"payment_methods"`
	OrderUpperLimit float64                  `json:"order_upper_limit"`
	OrderLowerLimit float64                  `json:"order_lower_limit"`
	ChainId         uint                     `json:"chain_id"`
	AssetType       reusable.AssetType       `json:"assetType"`
	Amount          float64                  `json:"amount"`
	Balance         float64                  `json:"balance"`
	CreatedAt       time.Time                `json:"created_at"`
	UpdatedAt       time.Time                `json:"updated_at"`
	Active          bool                     `json:"active"`
}
