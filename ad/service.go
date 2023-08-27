package ad

import (
	"errors"
)

type ServiceDbInterface interface {
	Create(StoreInputDto) (*Dto, error)
	Find(id string) (*Dto, error)
	UpdateActive(id string, active bool) (*Dto, error)
	List(ServiceListInputDto) ([]*Dto, error)
}

var (
	DbError = errors.New("db error")
)

type Service struct {
	db ServiceDbInterface
}

func (s Service) Store(input StoreInputDto) (*Dto, error) {
	if failed, err := input.Validate(); failed {
		return nil, err
	}
	ad, err := s.db.Create(input)
	if err != nil {
		return nil, err
	}

	return ad, nil
}

func (s Service) Find(id string) (*Dto, error) {
	ad, err := s.db.Find(id)
	if err != nil {
		return nil, DbError
	}

	return ad, nil
}

func (s Service) UpdateActive(id string, active bool) (*Dto, error) {
	ad, err := s.db.UpdateActive(id, active)
	if err != nil {
		return nil, DbError
	}

	return ad, nil
}

func (s Service) List(input ServiceListInputDto) ([]*Dto, error) {
	ads, err := s.db.List(input)
	if err != nil {
		return nil, err
	}

	return ads, nil
}

func (s Service) Search(inputDto ServiceListInputDto) ([]*Dto, error) {
	if failed, err := inputDto.Validate(); failed {
		return nil, err
	}

	return s.List(inputDto)
}

func NewService(db ServiceDbInterface) *Service {
	return &Service{db}
}

func NewMgmService() *Service {
	return NewService(NewServiceMdmAdapter())
}
