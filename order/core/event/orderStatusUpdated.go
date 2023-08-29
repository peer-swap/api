package event

import (
	"peerswap/order/core/dto"
	"peerswap/reusable"
)

type OrderStatusUpdated struct {
	Order  *dto.Order
	Status reusable.OrderStatus
}

func NewOrderStatusUpdated(order *dto.Order, status reusable.OrderStatus) *OrderStatusUpdated {
	return &OrderStatusUpdated{Order: order, Status: status}
}
