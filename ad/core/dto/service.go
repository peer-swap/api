package dto

import (
	"github.com/go-playground/validator"
	"peerswap/reusable"
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
