package core

import (
	"errors"
	"peerswap/order/core/dto"
	"peerswap/order/core/event"
	"peerswap/reusable"
)

type DbInterface interface {
	Create(*dto.ServiceStoreInput) (*dto.Order, error)
	FindAd(*dto.ServiceStoreInput) (*dto.Ad, error)
	DecrementAdBalance(string, float64) (*dto.Ad, error)
	UpdateStatus(string, reusable.OrderStatus) (*dto.Order, error)
	Cancel(string, *dto.ServiceCancelInput) (*dto.Order, error)
	Appeal(string, *dto.ServiceAppealInput) (*dto.Order, error)
	Find(id string)
}

type EscrowInterface interface {
	PlaceOrder(dto.PlaceOrderInput) error
	ReleaseToken(dto.PlaceOrderInput) error
	UnfreezeToken(dto.PlaceOrderInput) error
}

type Service struct {
	db      DbInterface
	escrow  EscrowInterface
	emitter reusable.Emitter
}

var (
	NotFoundError = errors.New("not found")
	CantUpdateErr = errors.New("cant update")
)

func (s Service) Store(input *dto.ServiceStoreInput) (*dto.Order, *dto.Ad, error) {
	if fails, err := reusable.NewValidator(input).Validate(); fails {
		return nil, nil, err
	}

	ad, err := s.db.FindAd(input)
	if err != nil {
		return nil, nil, err
	}

	order, err := s.db.Create(input)
	if err != nil {
		return nil, nil, err
	}

	err = s.escrow.PlaceOrder(dto.PlaceOrderInput{
		Ad:     input.Ad,
		Amount: input.Amount,
		Type:   input.Type,
		Asset:  input.Asset,
		Fiat:   input.Fiat,
		Seller: dto.PlaceOrderInputUser{Address: input.Seller.Address},
		Buyer:  dto.PlaceOrderInputUser{Address: input.Buyer.Address},
	})
	if err != nil {
		return nil, nil, err
	}
	ad, err = s.db.DecrementAdBalance(input.Ad, input.Amount)
	if err != nil {
		return nil, nil, err
	}

	s.emitter.Emit(event.NewOrderCreated(order, ad))

	return order, ad, nil
}

func (s Service) PaymentSent(id string) (*dto.Order, error) {
	order, err := s.db.UpdateStatus(id, reusable.OrderStatusPaymentSent)
	if err != nil {
		return nil, err
	}

	s.emitter.Emit(event.NewOrderStatusUpdated(order, reusable.OrderStatusPaymentSent))

	return order, nil
}

func (s Service) PaymentReceived(id string) (*dto.Order, error) {
	order, err := s.db.UpdateStatus(id, reusable.OrderStatusPaymentReceived)
	if err != nil {
		return nil, err
	}

	err = s.escrow.ReleaseToken(dto.PlaceOrderInput{
		Ad:     order.Ad,
		Amount: order.Amount,
		Type:   order.Type,
		Asset:  order.Asset,
		Fiat:   order.Fiat,
		Seller: dto.PlaceOrderInputUser{Address: order.Seller.Address},
		Buyer:  dto.PlaceOrderInputUser{Address: order.Buyer.Address},
	})
	if err != nil {
		return nil, err
	}

	s.emitter.Emit(event.NewOrderStatusUpdated(order, reusable.OrderStatusPaymentReceived))

	return order, nil
}

func (s Service) Cancel(id string, input *dto.ServiceCancelInput) (*dto.Order, error) {
	order, err := s.db.Cancel(id, input)
	if err != nil {
		return nil, err
	}

	err = s.escrow.UnfreezeToken(dto.PlaceOrderInput{
		Ad:     order.Ad,
		Amount: order.Amount,
		Type:   order.Type,
		Asset:  order.Asset,
		Fiat:   order.Fiat,
		Seller: dto.PlaceOrderInputUser{Address: order.Seller.Address},
		Buyer:  dto.PlaceOrderInputUser{Address: order.Buyer.Address},
	})
	if err != nil {
		return nil, err
	}

	s.emitter.Emit(event.NewOrderStatusUpdated(order, reusable.OrderStatusCanceled))

	return order, nil
}

func (s Service) Appeal(id string, input *dto.ServiceAppealInput) (*dto.Order, error) {
	order, err := s.db.Appeal(id, input)
	if err != nil {
		return nil, err
	}

	s.emitter.Emit(event.NewOrderStatusUpdated(order, reusable.OrderStatusAppealed))

	return order, nil
}

func NewService(db DbInterface, escrow EscrowInterface, emitter reusable.Emitter) *Service {
	return &Service{db: db, escrow: escrow, emitter: emitter}
}
