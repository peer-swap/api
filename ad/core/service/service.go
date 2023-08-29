package service

import (
	"errors"
	"peerswap/ad/core/dto"
	"peerswap/ad/core/event"
	"peerswap/reusable"
)

type DbInterface interface {
	DbFinderInterface
	Create(dto.StoreInputDto) (*dto.Ad, error)
	UpdateActive(id string, active bool) (*dto.Ad, error)
	List(dto.ServiceListInputDto) ([]*dto.Ad, error)
}

type DbFinderInterface interface {
	Find(id string) (*dto.Ad, error)
}

var (
	DbError = errors.New("db error")
)

type Service struct {
	db      DbInterface
	emitter reusable.Emitter
}

func (s Service) Store(input dto.StoreInputDto) (*dto.Ad, error) {
	if failed, err := input.Validate(); failed {
		return nil, err
	}
	ad, err := s.db.Create(input)
	if err != nil {
		return nil, err
	}

	s.emitter.Emit(event.AdCreated{Ad: ad})

	return ad, nil
}

func (s Service) Find(id string) (*dto.Ad, error) {
	ad, err := s.db.Find(id)
	if err != nil {
		return nil, DbError
	}

	return ad, nil
}

func (s Service) UpdateActive(id string, active bool) (*dto.Ad, error) {
	ad, err := s.db.UpdateActive(id, active)
	if err != nil {
		return nil, DbError
	}

	s.emitter.Emit(event.AdUpdatedActive{Ad: ad})

	return ad, nil
}

func (s Service) List(input dto.ServiceListInputDto) ([]*dto.Ad, error) {
	ads, err := s.db.List(input)
	if err != nil {
		return nil, err
	}

	return ads, nil
}

func (s Service) Search(inputDto dto.ServiceListInputDto) ([]*dto.Ad, error) {
	if failed, err := inputDto.Validate(); failed {
		return nil, err
	}

	return s.List(inputDto)
}

func NewService(db DbInterface, emitter reusable.Emitter) *Service {
	return &Service{db: db, emitter: emitter}
}
