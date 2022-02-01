package activationAuth

type InputActivation struct {
	ActivationCode string `json:"token" validate:"required"`
}
