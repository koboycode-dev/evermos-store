package route

import (
	createProduct "github.com/firmanJS/store-app/controllers/product-controllers/create"
	deleteProduct "github.com/firmanJS/store-app/controllers/product-controllers/delete"
	resultProduct "github.com/firmanJS/store-app/controllers/product-controllers/result"
	resultsProduct "github.com/firmanJS/store-app/controllers/product-controllers/results"
	updateProduct "github.com/firmanJS/store-app/controllers/product-controllers/update"
	handlerCreateProduct "github.com/firmanJS/store-app/handlers/product-handlers/create"
	handlerDeleteProduct "github.com/firmanJS/store-app/handlers/product-handlers/delete"
	handlerResultProduct "github.com/firmanJS/store-app/handlers/product-handlers/result"
	handlerResultsProduct "github.com/firmanJS/store-app/handlers/product-handlers/results"
	handlerUpdateProduct "github.com/firmanJS/store-app/handlers/product-handlers/update"
	middleware "github.com/firmanJS/store-app/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitProductRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Product
	*/
	createProductRepository := createProduct.NewRepositoryCreate(db)
	createProductService := createProduct.NewServiceCreate(createProductRepository)
	createProductHandler := handlerCreateProduct.NewHandlerCreateProduct(createProductService)

	resultsProductRepository := resultsProduct.NewRepositoryResults(db)
	resultsProductService := resultsProduct.NewServiceResults(resultsProductRepository)
	resultsProductHandler := handlerResultsProduct.NewHandlerResultsProduct(resultsProductService)

	resultProductRepository := resultProduct.NewRepositoryResult(db)
	resultProductService := resultProduct.NewServiceResult(resultProductRepository)
	resultProductHandler := handlerResultProduct.NewHandlerResultProduct(resultProductService)

	deleteProductRepository := deleteProduct.NewRepositoryDelete(db)
	deleteProductService := deleteProduct.NewServiceDelete(deleteProductRepository)
	deleteProductHandler := handlerDeleteProduct.NewHandlerDeleteProduct(deleteProductService)

	updateProductRepository := updateProduct.NewRepositoryUpdate(db)
	updateProductService := updateProduct.NewServiceUpdate(updateProductRepository)
	updateProductHandler := handlerUpdateProduct.NewHandlerUpdateProduct(updateProductService)

	/**
	@description All Product Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/product", createProductHandler.CreateProductHandler)
	groupRoute.GET("/product", resultsProductHandler.ResultsProductHandler)
	groupRoute.GET("/product/:id", resultProductHandler.ResultProductHandler)
	groupRoute.DELETE("/product/:id", deleteProductHandler.DeleteProductHandler)
	groupRoute.PUT("/product/:id", updateProductHandler.UpdateProductHandler)
}
