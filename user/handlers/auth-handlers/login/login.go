package handlerLogin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	loginAuth "github.com/svaqqosov/k8s_microserices_starter/controllers/auth-controllers/login"
	util "github.com/svaqqosov/k8s_microserices_starter/utils"
)

type handler struct {
	service loginAuth.Service
}

func NewHandlerLogin(service loginAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) LoginHandler(ctx *gin.Context) {

	var input loginAuth.InputLogin
	ctx.ShouldBindJSON(&input)

	resultLogin, errLogin := h.service.LoginService(&input)

	switch errLogin {

	case "LOGIN_NOT_FOUND_404":
		util.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodPost, nil)
		return

	case "LOGIN_NOT_ACTIVE_403":
		util.APIResponse(ctx, "User account is not active", http.StatusForbidden, http.MethodPost, nil)
		return

	case "LOGIN_WRONG_PASSWORD_403":
		util.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		jwtToken := util.NewJWT()
		accessTokenData := map[string]interface{}{"id": resultLogin.ID, "email": resultLogin.Email}
		accessToken, errToken := jwtToken.Create(24*60*1, accessTokenData)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		util.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, map[string]string{"accessToken": accessToken})
	}
}
