package escrow

import (
	"peerswap/order/core/dto"
)

type Escrow struct {
}

func (e Escrow) PlaceOrder(input dto.PlaceOrderInput) error {
	//TODO implement me
	return nil
}

func NewEscrow() *Escrow {
	return &Escrow{}
}
