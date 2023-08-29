package dto

type AdFundingAddTokenInput struct {
	ChainId       int     `json:"chain_id" validation:"required,numeric"`
	TransactionId string  `json:"transaction_id" validation:"required,string"`
	Amount        float64 `json:"amount"`
}

type AdFundingAddErc20Input struct {
	ChainId         int     `json:"chain_id" validation:"required,numeric"`
	MerchantAddress string  `json:"merchant_address" validation:"required"`
	TokenAddress    string  `json:"token_address" validation:"required"`
	Amount          float64 `json:"amount"`
}

type FundingVerificationResponse struct {
	Amount float64 `json:"amount"`
}
