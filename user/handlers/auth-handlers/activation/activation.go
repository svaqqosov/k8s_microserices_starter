package handlerActivation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	activationAuth "github.com/svaqqosov/k8s_microserices_starter/controllers/auth-controllers/activation"
	util "github.com/svaqqosov/k8s_microserices_starter/utils"
)

type handler struct {
	service activationAuth.Service
}

func NewHandlerActivation(service activationAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ActivationHandler(ctx *gin.Context) {
	var input activationAuth.InputActivation
	input.ActivationCode = ctx.Param("token")

	_, errActivation := h.service.ActivationService(&input)

	switch errActivation {

	case "ACTIVATION_NOT_FOUND_404":
		util.APIResponse(ctx, "User account is not exist", http.StatusNotFound, http.MethodPost, nil)
		return

	case "ACTIVATION_ACTIVE_400":
		util.APIResponse(ctx, "User account has been active please login", http.StatusBadRequest, http.MethodPost, nil)
		return

	case "ACTIVATION_ACCOUNT_FAILED_403":
		util.APIResponse(ctx, "Activation account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Activation account success", http.StatusOK, http.MethodPost, nil)
	}
}
