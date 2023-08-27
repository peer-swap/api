package ad

import (
	"github.com/go-playground/validator"
	"peerswap/reusable"
	"time"
)

type StoreInputDto struct {
	Type            reusable.TransactionType `json:"type" validate:"required"`
	Asset           string                   `json:"asset" validate:"required"`
	AssetType       reusable.AssetType       `json:"asset_type"`
	Fiat            string                   `json:"fiat" validate:"required"`
	Price           float64                  `json:"price" validate:"required"`
	Supply          float64                  `json:"supply" validate:"required"`
	PaymentMethods  []string                 `json:"payment_methods" validate:"required,dive,required"`
	OrderLowerLimit float64                  `json:"order_lower_limit" validate:"required,ltfield=OrderUpperLimit"`
	OrderUpperLimit float64                  `json:"order_upper_limit" validate:"required,gtfield=OrderLowerLimit"`
	ChainId         uint                     `json:"chain_id" validate:"required"`
}

func (d StoreInputDto) Validate() (bool, error) {
	validate := validator.New()
	err := validate.Struct(d)
	return err != nil, err
}

type PaymentMethodField struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	ContentType string `json:"content_type"`
	IsRequired  bool   `json:"is_required"`
	IsCopyable  bool   `json:"is_copyable"`
	IsDisplay   bool   `json:"is_display"`
	Value       string `json:"value"`
}

type PaymentMethod struct {
	MethodId string               `json:"methodId"`
	Display  string               `json:"display"`
	Method   string               `json:"method"`
	Fields   []PaymentMethodField `json:"fields,omitempty"`
}

type Dto struct {
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

type ServiceListInputDto struct {
	Type           reusable.TransactionType `validate:"required,alpha"`
	Amount         float64                  `validate:"numeric"`
	Fiat           string                   `validate:"required,alphanum"`
	Asset          string                   `validate:"required,alphanum"`
	PaymentMethods []string
	ChainId        uint
}

func (s ServiceListInputDto) Validate() (bool, *reusable.ValidationErrorsMapper) {
	validate := validator.New()
	err := validate.Struct(s)

	if err != nil {
		return true, &reusable.ValidationErrorsMapper{ValidationErrors: err.(validator.ValidationErrors)}
	}
	return false, nil
}
