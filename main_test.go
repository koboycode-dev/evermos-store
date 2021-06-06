package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	resultsProduct "github.com/firmanJS/store-app/controllers/product-controllers/results"
	handlerResultsProduct "github.com/firmanJS/store-app/handlers/product-handlers/results"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestControllerGetAll(t *testing.T) {

	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)
	var db *gorm.DB
	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	resultsProductRepository := resultsProduct.NewRepositoryResults(db)
	resultsProductService := resultsProduct.NewServiceResults(resultsProductRepository)
	resultsProductHandler := handlerResultsProduct.NewHandlerResultsProduct(resultsProductService)
	r.GET("/api/v1/product", resultsProductHandler.ResultsProductHandler)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusNotFound {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
