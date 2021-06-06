package handlerDeleteProduct

import (
	"net/http"

	deleteProduct "github.com/firmanJS/store-app/controllers/product-controllers/delete"
	util "github.com/firmanJS/store-app/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service deleteProduct.Service
}

func NewHandlerDeleteProduct(service deleteProduct.Service) *handler {
	return &handler{service: service}
}

func (h *handler) DeleteProductHandler(ctx *gin.Context) {

	var input deleteProduct.InputDeleteProduct
	input.Id = ctx.Param("id")

	_, errDeleteProduct := h.service.DeleteProductService(&input)

	switch errDeleteProduct {

	case util.NOT_FOUND:
		util.APIResponse(ctx, "Product data is not exist or deleted", http.StatusForbidden, http.MethodDelete, nil)
		return

	case util.FAILED:
		util.APIResponse(ctx, "Delete Product data failed", http.StatusForbidden, http.MethodDelete, nil)
		return

	default:
		util.APIResponse(ctx, input.Id, http.StatusOK, http.MethodDelete, nil)
	}
}
