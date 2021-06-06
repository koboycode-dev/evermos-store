package resultCart

import (
	model "github.com/firmanJS/store-app/models"
)

type Service interface {
	ResultCartService(input *InputResultCart) (*model.EntityCart, string)
}

type service struct {
	repository Repository
}

func NewServiceResult(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResultCartService(input *InputResultCart) (*model.EntityCart, string) {

	Carts := model.EntityCart{
		Id: input.Id,
	}

	resultGetCart, errGetCart := s.repository.ResultCartRepository(&Carts)

	return resultGetCart, errGetCart
}
