package handlerCreateCart

import (
	"net/http"
	"strings"

	createCart "github.com/firmanJS/store-app/controllers/cart-controllers/create"
	util "github.com/firmanJS/store-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	service createCart.Service
}

func NewHandlerCreateCart(service createCart.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateCartHandler(ctx *gin.Context) {

	var input createCart.InputCreateCart
	reqToken := ctx.GetHeader("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	resultToken, _ := util.VerifyToken(reqToken, "JWT_SECRET")

	result := util.DecodeToken(resultToken)

	input.User_Id = result.Claims.ID
	ctx.ShouldBindJSON(&input)
	if err := validator.New().Struct(input); err != nil {
		util.APIResponse(ctx, err.Error(), http.StatusBadRequest, http.MethodPost, nil)
	} else {

		_, errCreateCart := h.service.CreateCartService(&input)

		switch errCreateCart {

		case util.CONFLICT:
			util.APIResponse(ctx, "ID Cart already exist", http.StatusConflict, http.MethodPost, nil)
			return

		case util.FAILED:
			util.APIResponse(ctx, "Create new Cart failed", http.StatusForbidden, http.MethodPost, nil)
			return

		default:
			util.APIResponse(ctx, "Create new Cart successfully", http.StatusCreated, http.MethodPost, input)
		}
	}
}
