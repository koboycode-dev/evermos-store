package handlerResultProduct

import (
	"net/http"

	resultProduct "github.com/firmanJS/store-app/controllers/product-controllers/result"
	util "github.com/firmanJS/store-app/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service resultProduct.Service
}

func NewHandlerResultProduct(service resultProduct.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResultProductHandler(ctx *gin.Context) {

	var input resultProduct.InputResultProduct
	input.Id = ctx.Param("id")

	resultProduct, errResultProduct := h.service.ResultProductService(&input)

	switch errResultProduct {

	case util.NOT_FOUND:
		util.APIResponse(ctx, "Product data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		util.APIResponse(ctx, "Get Product data successfully", http.StatusOK, http.MethodGet, resultProduct)
	}
}
