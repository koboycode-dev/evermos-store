package handlerDeleteCart

import (
	"net/http"

	deleteCart "github.com/firmanJS/store-app/controllers/cart-controllers/delete"
	util "github.com/firmanJS/store-app/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service deleteCart.Service
}

func NewHandlerDeleteCart(service deleteCart.Service) *handler {
	return &handler{service: service}
}

func (h *handler) DeleteCartHandler(ctx *gin.Context) {

	var input deleteCart.InputDeleteCart
	input.Id = ctx.Param("id")

	_, errDeleteCart := h.service.DeleteCartService(&input)

	switch errDeleteCart {

	case util.NOT_FOUND:
		util.APIResponse(ctx, "Cart data is not exist or deleted", http.StatusForbidden, http.MethodDelete, nil)
		return

	case util.FAILED:
		util.APIResponse(ctx, "Delete Cart data failed", http.StatusForbidden, http.MethodDelete, nil)
		return

	default:
		util.APIResponse(ctx, input.Id, http.StatusOK, http.MethodDelete, nil)
	}
}
