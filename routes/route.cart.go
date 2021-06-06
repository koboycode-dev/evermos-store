package route

import (
	createCart "github.com/firmanJS/store-app/controllers/cart-controllers/create"
	deleteCart "github.com/firmanJS/store-app/controllers/cart-controllers/delete"
	resultCart "github.com/firmanJS/store-app/controllers/cart-controllers/result"
	resultsCart "github.com/firmanJS/store-app/controllers/cart-controllers/results"
	updateCart "github.com/firmanJS/store-app/controllers/cart-controllers/update"
	handlerCreateCart "github.com/firmanJS/store-app/handlers/cart-handlers/create"
	handlerDeleteCart "github.com/firmanJS/store-app/handlers/cart-handlers/delete"
	handlerResultCart "github.com/firmanJS/store-app/handlers/cart-handlers/result"
	handlerResultsCart "github.com/firmanJS/store-app/handlers/cart-handlers/results"
	handlerUpdateCart "github.com/firmanJS/store-app/handlers/cart-handlers/update"
	middleware "github.com/firmanJS/store-app/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitCartRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Cart
	*/
	createCartRepository := createCart.NewRepositoryCreate(db)
	createCartService := createCart.NewServiceCreate(createCartRepository)
	createCartHandler := handlerCreateCart.NewHandlerCreateCart(createCartService)

	resultsCartRepository := resultsCart.NewRepositoryResults(db)
	resultsCartService := resultsCart.NewServiceResults(resultsCartRepository)
	resultsCartHandler := handlerResultsCart.NewHandlerResultsCart(resultsCartService)

	resultCartRepository := resultCart.NewRepositoryResult(db)
	resultCartService := resultCart.NewServiceResult(resultCartRepository)
	resultCartHandler := handlerResultCart.NewHandlerResultCart(resultCartService)

	deleteCartRepository := deleteCart.NewRepositoryDelete(db)
	deleteCartService := deleteCart.NewServiceDelete(deleteCartRepository)
	deleteCartHandler := handlerDeleteCart.NewHandlerDeleteCart(deleteCartService)

	updateCartRepository := updateCart.NewRepositoryUpdate(db)
	updateCartService := updateCart.NewServiceUpdate(updateCartRepository)
	updateCartHandler := handlerUpdateCart.NewHandlerUpdateCart(updateCartService)

	/**
	@description All Cart Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/cart", createCartHandler.CreateCartHandler)
	groupRoute.GET("/cart", resultsCartHandler.ResultsCartHandler)
	groupRoute.GET("/cart/:id", resultCartHandler.ResultCartHandler)
	groupRoute.DELETE("/cart/:id", deleteCartHandler.DeleteCartHandler)
	groupRoute.PUT("/cart/:id", updateCartHandler.UpdateCartHandler)
}
