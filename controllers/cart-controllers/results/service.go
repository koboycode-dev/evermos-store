package resultsCart

import (
	model "github.com/firmanJS/store-app/models"
)

type Service interface {
	ResultsCartService() (*[]model.EntityCart, string)
}

type service struct {
	repository Repository
}

func NewServiceResults(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResultsCartService() (*[]model.EntityCart, string) {

	resultCart, errCart := s.repository.ResultsCartRepository()

	return resultCart, errCart
}
