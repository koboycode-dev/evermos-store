package resultsProduct

import (
	model "github.com/firmanJS/store-app/models"
)

type Service interface {
	ResultsProductService() (*[]model.EntityProduct, string)
}

type service struct {
	repository Repository
}

func NewServiceResults(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResultsProductService() (*[]model.EntityProduct, string) {

	resultProduct, errProduct := s.repository.ResultsProductRepository()

	return resultProduct, errProduct
}
