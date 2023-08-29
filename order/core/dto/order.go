package dto

import (
	"peerswap/reusable"
	"time"
)

type User struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Mobile   string `json:"mobile"`
	Address  string `json:"address" validate:"required,alphanum"`
}
type Canceled struct {
	At     time.Time `json:"at"`
	By     string    `json:"by"`
	Reason string    `json:"reason"`
}
type Appealed struct {
	At          time.Time `json:"at"`
	By          string    `json:"by"`
	Reason      string    `json:"reason"`
	Attachments []string  `json:"attachments"`
}

type Order struct {
	Id             string                   `json:"id"`
	Ad             string                   `json:"ad"`
	Amount         float64                  `json:"amount"`
	Type           reusable.TransactionType `json:"type"`
	Asset          string                   `json:"asset"`
	Fiat           string                   `json:"fiat"`
	MatchPrice     float64                  `json:"matchPrice"`
	Seller         User                     `json:"seller,omitempty"`
	Buyer          User                     `json:"buyer,omitempty"`
	MerchantId     string                   `json:"merchantId"`
	Canceled       Canceled                 `json:"canceled,omitempty"`
	Appealed       Appealed                 `json:"appealed,omitempty"`
	CreatedAt      time.Time                `json:"createdAt"`
	UpdatedAt      time.Time                `json:"updatedAt"`
	NotifiedPaidAt time.Time                `json:"notifiedPaidAt"`
	ConfirmPaidAt  time.Time                `json:"confirmPaidAt"`
	ReleasedAt     time.Time                `json:"releasedAt"`
	Status         reusable.OrderStatus     `json:"status"`
}
