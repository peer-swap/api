package ad

import (
	"github.com/kamva/mgm/v3"
	"peerswap/reusable"
	"time"
)

type Ad struct {
	mgm.DefaultModel `bson:",inline"`
	Type             reusable.TransactionType `bson:"type"`
	Asset            string                   `bson:"asset"`
	Fiat             string                   `bson:"fiat"`
	Price            float64                  `bson:"price"`
	Supply           float64                  `bson:"supply"`
	PaymentMethods   []ModelPaymentMethod     `bson:"paymentMethods"`
	OrderUpperLimit  float64                  `bson:"orderUpperLimit"`
	OrderLowerLimit  float64                  `bson:"orderLowerLimit"`
	ChainId          uint                     `bson:"chainId"`
	Merchant         Merchant                 `bson:"merchant,omitempty"`
	AssetType        reusable.AssetType       `bson:"assetType"`
	Amount           float64                  `bson:"amount"`
	Balance          float64                  `bson:"balance"`
	Status           reusable.AdStatus        `bson:"status"`
	StoppedAt        time.Time                `bson:"stoppedAt"`
	Active           bool                     `bson:"active"`
}

func NewModelFromStoreInputDto(input StoreInputDto) *Ad {
	return &Ad{
		Type:            input.Type,
		Asset:           input.Asset,
		Fiat:            input.Fiat,
		Price:           input.Price,
		Supply:          input.Supply,
		PaymentMethods:  []ModelPaymentMethod{},
		OrderUpperLimit: input.OrderUpperLimit,
		OrderLowerLimit: input.OrderLowerLimit,
		ChainId:         input.ChainId,
		AssetType:       input.AssetType,
	}
}

func (m *Ad) ToDto() *Dto {
	var methods []PaymentMethod

	for _, method := range m.PaymentMethods {
		var fields []PaymentMethodField

		for _, field := range method.Fields {
			fields = append(fields, PaymentMethodField{
				Id:          field.Id,
				Name:        field.Name,
				ContentType: field.ContentType,
				IsRequired:  field.IsRequired,
				IsCopyable:  field.IsCopyable,
				IsDisplay:   field.IsDisplay,
				Value:       field.Value,
			})
		}
		methods = append(methods, PaymentMethod{
			MethodId: method.MethodId,
			Display:  method.Display,
			Method:   method.Method,
			Fields:   fields,
		})
	}

	return &Dto{
		Type:            m.Type,
		Asset:           m.Asset,
		Fiat:            m.Fiat,
		Price:           m.Price,
		Amount:          m.Amount,
		Balance:         m.Balance,
		PaymentMethods:  methods,
		OrderUpperLimit: m.OrderUpperLimit,
		OrderLowerLimit: m.OrderLowerLimit,
		ChainId:         m.ChainId,
		AssetType:       m.AssetType,
		Id:              m.ID.Hex(),
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		Active:          m.Active,
	}
}
