package handlerCreateProduct

import (
	"net/http"

	createProduct "github.com/firmanJS/store-app/controllers/product-controllers/create"
	util "github.com/firmanJS/store-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	service createProduct.Service
}

func NewHandlerCreateProduct(service createProduct.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateProductHandler(ctx *gin.Context) {

	var input createProduct.InputCreateProduct
	ctx.ShouldBindJSON(&input)

	if err := validator.New().Struct(input); err != nil {
		util.APIResponse(ctx, err.Error(), http.StatusBadRequest, http.MethodPost, nil)
	} else {

		_, errCreateProduct := h.service.CreateProductService(&input)

		switch errCreateProduct {

		case util.CONFLICT:
			util.APIResponse(ctx, "ID Product already exist", http.StatusConflict, http.MethodPost, nil)
			return

		case util.FAILED:
			util.APIResponse(ctx, "Create new Product failed", http.StatusForbidden, http.MethodPost, nil)
			return

		default:
			util.APIResponse(ctx, "Create new Product successfully", http.StatusCreated, http.MethodPost, input)
		}
	}
}
