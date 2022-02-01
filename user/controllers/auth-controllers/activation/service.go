package activationAuth

import model "github.com/svaqqosov/k8s_microserices_starter/models"

type Service interface {
	ActivationService(input *InputActivation) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceActivation(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ActivationService(input *InputActivation) (*model.EntityUsers, string) {
	users := model.EntityUsers{
		ActivationCode: input.ActivationCode,
	}

	activationResult, activationError := s.repository.ActivationRepository(&users)

	return activationResult, activationError
}
