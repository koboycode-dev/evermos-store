package handlerResultsCart

import (
	"net/http"

	resultsCart "github.com/firmanJS/store-app/controllers/cart-controllers/results"
	util "github.com/firmanJS/store-app/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service resultsCart.Service
}

func NewHandlerResultsCart(service resultsCart.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResultsCartHandler(ctx *gin.Context) {

	resultsCart, errResultsCart := h.service.ResultsCartService()

	switch errResultsCart {

	case util.NOT_FOUND:
		util.APIResponse(ctx, "Carts data is not exists", http.StatusConflict, http.MethodGet, nil)

	default:
		util.APIResponse(ctx, "Get Carts data successfully", http.StatusOK, http.MethodGet, resultsCart)
	}
}
