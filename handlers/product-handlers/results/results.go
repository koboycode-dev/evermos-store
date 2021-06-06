package handlerResultsProduct

import (
	"net/http"

	resultsProduct "github.com/firmanJS/store-app/controllers/product-controllers/results"
	util "github.com/firmanJS/store-app/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service resultsProduct.Service
}

func NewHandlerResultsProduct(service resultsProduct.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResultsProductHandler(ctx *gin.Context) {

	resultsProduct, errResultsProduct := h.service.ResultsProductService()

	switch errResultsProduct {

	case util.NOT_FOUND:
		util.APIResponse(ctx, "Products data is not exists", http.StatusConflict, http.MethodGet, nil)

	default:
		util.APIResponse(ctx, "Get Products data successfully", http.StatusOK, http.MethodGet, resultsProduct)
	}
}
