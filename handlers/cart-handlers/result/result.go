package handlerResultCart

import (
	"net/http"

	resultCart "github.com/firmanJS/store-app/controllers/cart-controllers/result"
	util "github.com/firmanJS/store-app/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service resultCart.Service
}

func NewHandlerResultCart(service resultCart.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResultCartHandler(ctx *gin.Context) {

	var input resultCart.InputResultCart
	input.Id = ctx.Param("id")

	resultCart, errResultCart := h.service.ResultCartService(&input)

	switch errResultCart {

	case util.NOT_FOUND:
		util.APIResponse(ctx, "Cart data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		util.APIResponse(ctx, "Get Cart data successfully", http.StatusOK, http.MethodGet, resultCart)
	}
}
