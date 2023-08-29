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
	Supply          float64                  `json:"supply"`
	OrderUpperLimit float64                  `json:"orderUpperLimit"`
	OrderLowerLimit float64                  `json:"orderLowerLimit"`
	ChainId         uint                     `json:"chainId"`
	Balance         float64                  `json:"balance"`
	Status          reusable.AdStatus        `json:"status"`
	CreatedAt       time.Time                `json:"createdAt"`
	UpdatedAt       time.Time                `json:"updatedAt"`
}
