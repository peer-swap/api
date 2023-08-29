package mongo

import (
	"github.com/kamva/mgm/v3"
	dto2 "peerswap/order/core/dto"
	"peerswap/reusable"
	"time"
)

type User struct {
	Name     string `bson:"name"`
	Nickname string `bson:"nickname"`
	Mobile   string `bson:"mobile"`
	Address  string `bson:"address"`
}
type Canceled struct {
	At     time.Time `bson:"at"`
	By     string    `bson:"by"`
	Reason string    `bson:"reason"`
}
type Appealed struct {
	At          time.Time `bson:"at"`
	By          string    `bson:"by"`
	Reason      string    `bson:"reason"`
	Attachments []string  `bson:"attachments"`
}

type Order struct {
	mgm.DefaultModel `bson:",inline"`
	Ad               string                   `bson:"ad"`
	Amount           float64                  `bson:"amount"`
	Type             reusable.TransactionType `bson:"type"`
	Asset            string                   `bson:"asset"`
	Fiat             string                   `bson:"fiat"`
	MatchPrice       float64                  `bson:"matchPrice"`
	Seller           User                     `bson:"seller,omitempty"`
	Buyer            User                     `bson:"buyer,omitempty"`
	MerchantId       string                   `bson:"merchantId"`
	Canceled         Canceled                 `bson:"canceled,omitempty"`
	Appealed         Appealed                 `bson:"appealed,omitempty"`
	NotifiedPaidAt   time.Time                `bson:"notifiedPaidAt"`
	ConfirmPaidAt    time.Time                `bson:"confirmPaidAt"`
	ReleasedAt       time.Time                `bson:"releasedAt"`
	Status           reusable.OrderStatus     `json:"status"`
}

func NewOrderFromDtoStoreInput(props *dto2.ServiceStoreInput) *Order {
	return &Order{
		Ad:         props.Ad,
		Amount:     props.Amount,
		Type:       props.Type,
		Asset:      props.Asset,
		Fiat:       props.Fiat,
		MatchPrice: props.MatchPrice,
		Seller: User{
			Address: props.Seller.Address,
		},
		Buyer: User{
			Address: props.Buyer.Address,
		},
		MerchantId: props.MerchantId,
	}
}

func (o Order) ToDtoOrder() *dto2.Order {
	return &dto2.Order{
		Id:         o.ID.Hex(),
		Ad:         o.Ad,
		Amount:     o.Amount,
		Type:       o.Type,
		Asset:      o.Asset,
		Fiat:       o.Fiat,
		MatchPrice: o.MatchPrice,
		Seller: dto2.User{
			Name:    o.Seller.Name,
			Address: o.Seller.Address,
		},
		Buyer: dto2.User{
			Name:    o.Buyer.Name,
			Address: o.Buyer.Address,
		},
		MerchantId: o.MerchantId,
		Canceled: dto2.Canceled{
			At:     o.Appealed.At,
			By:     o.Appealed.By,
			Reason: o.Appealed.Reason,
		},
		Appealed: dto2.Appealed{
			At:          o.Appealed.At,
			By:          o.Appealed.By,
			Reason:      o.Appealed.Reason,
			Attachments: o.Appealed.Attachments,
		},
		CreatedAt:      o.CreatedAt,
		UpdatedAt:      o.UpdatedAt,
		NotifiedPaidAt: o.NotifiedPaidAt,
		ConfirmPaidAt:  o.ConfirmPaidAt,
		ReleasedAt:     o.ReleasedAt,
		Status:         o.Status,
	}
}
