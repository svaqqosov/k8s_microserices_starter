package handlerRegister

import (
	"net/http"

	"github.com/gin-gonic/gin"
	registerAuth "github.com/svaqqosov/k8s_microserices_starter/controllers/auth-controllers/register"
	util "github.com/svaqqosov/k8s_microserices_starter/utils"
)

type handler struct {
	service registerAuth.Service
}

func NewHandlerRegister(service registerAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {

	var input registerAuth.InputRegister
	ctx.ShouldBindJSON(&input)

	resultRegister, errRegister := h.service.RegisterService(&input)

	switch errRegister {

	case "REGISTER_CONFLICT_409":
		util.APIResponse(ctx, "Email already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "REGISTER_FAILED_403":
		util.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, resultRegister.ActivationCode)
	}
}
