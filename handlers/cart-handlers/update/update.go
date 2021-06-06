package handlerUpdateCart

import (
	"net/http"

	updateCart "github.com/firmanJS/store-app/controllers/cart-controllers/update"
	util "github.com/firmanJS/store-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	service updateCart.Service
}

func NewHandlerUpdateCart(service updateCart.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateCartHandler(ctx *gin.Context) {

	var input updateCart.InputUpdateCart
	input.Id = ctx.Param("id")
	ctx.ShouldBindJSON(&input)

	if err := validator.New().Struct(input); err != nil {
		util.APIResponse(ctx, err.Error(), http.StatusBadRequest, http.MethodPost, nil)
	} else {

		_, errUpdateCart := h.service.UpdateCartService(&input)

		switch errUpdateCart {

		case util.NOT_FOUND:
			util.APIResponse(ctx, "Cart data is not exist or deleted", http.StatusNotFound, http.MethodPost, nil)

		case util.FAILED:
			util.APIResponse(ctx, "Update Cart data failed", http.StatusForbidden, http.MethodPost, nil)

		default:
			util.APIResponse(ctx, "Update Cart data sucessfully", http.StatusOK, http.MethodPost, input)
		}

	}
}
