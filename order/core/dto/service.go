package dto

import "peerswap/reusable"

type ServiceStoreInput struct {
	Ad         string                   `json:"ad" validate:"required,alphanum"`
	Amount     float64                  `json:"amount" validate:"required,numeric"`
	Type       reusable.TransactionType `json:"type" validate:"required,alpha,oneof=BUY SELL"`
	Asset      string                   `json:"asset" validate:"required,alphanum"`
	Fiat       string                   `json:"fiat" validate:"required,alphanum"`
	MatchPrice float64                  `json:"matchPrice" validate:"required,numeric"`
	Seller     User                     `json:"seller,omitempty"`
	Buyer      User                     `json:"buyer,omitempty"`
	MerchantId string                   `json:"merchantId"`
}

type ServiceCancelInput struct {
	By     string `json:"by"`
	Reason string `json:"reason"`
}

type ServiceAppealInput struct {
	By          string   `json:"by"`
	Reason      string   `json:"reason"`
	Attachments []string `json:"attachments"`
}
