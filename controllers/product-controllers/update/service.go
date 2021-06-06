package updateProduct

import (
	model "github.com/firmanJS/store-app/models"
)

type Service interface {
	UpdateProductService(input *InputUpdateProduct) (*model.EntityProduct, string)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateProductService(input *InputUpdateProduct) (*model.EntityProduct, string) {

	Products := model.EntityProduct{
		Id:       input.Id,
		Name:     input.Name,
		Quantity: input.Quantity,
		Price:    input.Price,
	}

	resultUpdateProduct, errUpdateProduct := s.repository.UpdateProductRepository(&Products)

	return resultUpdateProduct, errUpdateProduct
}
