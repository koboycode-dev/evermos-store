package createCart

import (
	model "github.com/firmanJS/store-app/models"
)

type Service interface {
	CreateCartService(input *InputCreateCart) (*model.EntityCart, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateCartService(input *InputCreateCart) (*model.EntityCart, string) {

	Carts := model.EntityCart{
		Order_Id:   input.Order_Id,
		Product_Id: input.Product_Id,
		Quantity:   input.Quantity,
		Note:       input.Note,
	}

	resultCreateCart, errCreateCart := s.repository.CreateCartRepository(&Carts)

	return resultCreateCart, errCreateCart
}
