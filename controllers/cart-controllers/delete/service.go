package deleteCart

import (
	model "github.com/firmanJS/store-app/models"
)

type Service interface {
	DeleteCartService(input *InputDeleteCart) (*model.EntityCart, string)
}

type service struct {
	repository Repository
}

func NewServiceDelete(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) DeleteCartService(input *InputDeleteCart) (*model.EntityCart, string) {

	Carts := model.EntityCart{
		Id: input.Id,
	}

	resultDeleteCart, errDeleteCart := s.repository.DeletedCartRepository(&Carts)

	return resultDeleteCart, errDeleteCart
}
