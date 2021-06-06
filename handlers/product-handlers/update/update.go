package handlerUpdateProduct

import (
	"net/http"

	updateProduct "github.com/firmanJS/store-app/controllers/product-controllers/update"
	util "github.com/firmanJS/store-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	service updateProduct.Service
}

func NewHandlerUpdateProduct(service updateProduct.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateProductHandler(ctx *gin.Context) {

	var input updateProduct.InputUpdateProduct
	input.Id = ctx.Param("id")
	ctx.ShouldBindJSON(&input)

	if err := validator.New().Struct(input); err != nil {
		util.APIResponse(ctx, err.Error(), http.StatusBadRequest, http.MethodPost, nil)
	} else {

		_, errUpdateProduct := h.service.UpdateProductService(&input)

		switch errUpdateProduct {

		case util.NOT_FOUND:
			util.APIResponse(ctx, "Product data is not exist or deleted", http.StatusNotFound, http.MethodPut, nil)

		case util.FAILED:
			util.APIResponse(ctx, "Update Product data failed", http.StatusForbidden, http.MethodPut, nil)

		default:
			util.APIResponse(ctx, "Update Product data sucessfully", http.StatusOK, http.MethodPut, input)
		}

	}
}
