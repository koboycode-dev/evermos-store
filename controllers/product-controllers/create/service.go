package createProduct

import (
	model "github.com/firmanJS/store-app/models"
)

type Service interface {
	CreateProductService(input *InputCreateProduct) (*model.EntityProduct, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateProductService(input *InputCreateProduct) (*model.EntityProduct, string) {

	products := model.EntityProduct{
		Id_Product: input.Id_Product,
		Name:       input.Name,
		Price:      input.Price,
		Quantity:   input.Quantity,
	}

	resultCreateProduct, errCreateProduct := s.repository.CreateProductRepository(&products)

	return resultCreateProduct, errCreateProduct
}
