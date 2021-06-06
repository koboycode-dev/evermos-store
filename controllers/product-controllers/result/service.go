package resultProduct

import (
	model "github.com/firmanJS/store-app/models"
)

type Service interface {
	ResultProductService(input *InputResultProduct) (*model.EntityProduct, string)
}

type service struct {
	repository Repository
}

func NewServiceResult(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResultProductService(input *InputResultProduct) (*model.EntityProduct, string) {

	Products := model.EntityProduct{
		Id: input.Id,
	}

	resultGetProduct, errGetProduct := s.repository.ResultProductRepository(&Products)

	return resultGetProduct, errGetProduct
}
