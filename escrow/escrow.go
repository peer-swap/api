package escrow

import (
	"peerswap/order/dto"
)

type Escrow struct {
}

func (e Escrow) PlaceOrder(order *dto.ServiceStoreInput) error {
	return nil
}

func (e Escrow) ConfirmTransactionAndAmount(dto FundControllerAddTokenInputDto) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (e Escrow) ConfirmAllowanceMatchAmount(dto FundControllerAddCoinInputDto) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (e Escrow) DepositCoin(dto DepositCoinInputDto) error {
	//TODO implement me
	panic("implement me")
}

func NewEscrow() *Escrow {
	return &Escrow{}
}

type FundControllerAddTokenInputDto struct {
	ChainId       int     `json:"chain_id"`
	TransactionId string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
}

type FundControllerAddCoinInputDto struct {
	ChainId         int     `json:"chain_id"`
	MerchantAddress string  `json:"merchant_address"`
	TokenAddress    string  `json:"token_address"`
	Amount          float64 `json:"amount"`
}

type DepositCoinInputDto struct {
	ChainId         int     `json:"chain_id"`
	MerchantAddress string  `json:"merchant_address"`
	TokenAddress    string  `json:"token_address"`
	Amount          float64 `json:"amount"`
}
