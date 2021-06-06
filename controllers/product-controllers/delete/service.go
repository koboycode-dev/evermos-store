package deleteProduct

import (
	model "github.com/firmanJS/store-app/models"
)

type Service interface {
	DeleteProductService(input *InputDeleteProduct) (*model.EntityProduct, string)
}

type service struct {
	repository Repository
}

func NewServiceDelete(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) DeleteProductService(input *InputDeleteProduct) (*model.EntityProduct, string) {

	Products := model.EntityProduct{
		Id: input.Id,
	}

	resultDeleteProduct, errDeleteProduct := s.repository.DeletedProductRepository(&Products)

	return resultDeleteProduct, errDeleteProduct
}
