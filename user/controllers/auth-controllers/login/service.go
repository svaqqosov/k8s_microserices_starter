package loginAuth

import (
	model "github.com/svaqqosov/k8s_microserices_starter/models"
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
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.LoginRepository(&user)

	return resultLogin, errLogin
}
