package route

import (
	loginAuth "github.com/firmanJS/store-app/controllers/auth-controllers/login"
	registerAuth "github.com/firmanJS/store-app/controllers/auth-controllers/register"
	handlerLogin "github.com/firmanJS/store-app/handlers/auth-handlers/login"
	handlerRegister "github.com/firmanJS/store-app/handlers/auth-handlers/register"
	"github.com/gin-gonic/gin"
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

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)

}
