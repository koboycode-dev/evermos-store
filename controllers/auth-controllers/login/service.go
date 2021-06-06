package loginAuth

import (
	model "github.com/firmanJS/store-app/models"
)

type Service interface {
	LoginService(input *InputLogin) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginService(input *InputLogin) (*model.EntityUsers, string) {

	user := model.EntityUsers{
		Username: input.Username,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.LoginRepository(&user)

	return resultLogin, errLogin
}
