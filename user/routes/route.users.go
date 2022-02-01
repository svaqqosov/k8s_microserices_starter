package route

import (
	"github.com/gin-gonic/gin"
	activationAuth "github.com/svaqqosov/k8s_microserices_starter/controllers/auth-controllers/activation"
	loginAuth "github.com/svaqqosov/k8s_microserices_starter/controllers/auth-controllers/login"
	registerAuth "github.com/svaqqosov/k8s_microserices_starter/controllers/auth-controllers/register"
	handlerActivation "github.com/svaqqosov/k8s_microserices_starter/handlers/auth-handlers/activation"
	handlerLogin "github.com/svaqqosov/k8s_microserices_starter/handlers/auth-handlers/login"
	handlerRegister "github.com/svaqqosov/k8s_microserices_starter/handlers/auth-handlers/register"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Auth
	*/
	LoginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(LoginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)

	registerRepository := registerAuth.NewRepositoryRegister(db)
	registerService := registerAuth.NewServiceRegister(registerRepository)
	registerHandler := handlerRegister.NewHandlerRegister(registerService)

	activationRepository := activationAuth.NewRepositoryActivation(db)
	activationService := activationAuth.NewServiceActivation(activationRepository)
	activationHandler := handlerActivation.NewHandlerActivation(activationService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1/users")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)
	groupRoute.POST("/activation/:token", activationHandler.ActivationHandler)

}
