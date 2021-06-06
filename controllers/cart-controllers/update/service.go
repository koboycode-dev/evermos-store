package updateCart

import (
	model "github.com/firmanJS/store-app/models"
)

type Service interface {
	UpdateCartService(input *InputUpdateCart) (*model.EntityCart, string)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateCartService(input *InputUpdateCart) (*model.EntityCart, string) {

	Carts := model.EntityCart{
		Id:         input.Id,
		Product_Id: input.Product_Id,
		Quantity:   input.Quantity,
		Note:       input.Note,
	}

	resultUpdateCart, errUpdateCart := s.repository.UpdateCartRepository(&Carts)

	return resultUpdateCart, errUpdateCart
}
