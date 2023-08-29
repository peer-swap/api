package event

import (
	dto2 "peerswap/order/core/dto"
)

type OrderCreated struct {
	Order *dto2.Order
	Ad    *dto2.Ad
}

func NewOrderCreated(order *dto2.Order, status *dto2.Ad) *OrderCreated {
	return &OrderCreated{Order: order, Ad: status}
}
