package dto

import "time"

type Fiat struct {
	Id          string    `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	CountryCode string    `json:"country_code"`
	Symbol      string    `json:"symbol"`
	Icon        string    `json:"icon"`
	CratedAt    time.Time `json:"crated_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FiatListFilter struct {
	Query string `json:"query"`
}

type FiatAddInput struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	Symbol      string `json:"symbol"`
	Icon        string `json:"icon"`
}
