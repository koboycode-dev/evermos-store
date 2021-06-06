package handlerRegister

import (
	"net/http"

	registerAuth "github.com/firmanJS/store-app/controllers/auth-controllers/register"
	util "github.com/firmanJS/store-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service registerAuth.Service
}

func NewHandlerRegister(service registerAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {

	var input registerAuth.InputRegister
	ctx.ShouldBindJSON(&input)

	if err := validator.New().Struct(input); err != nil {
		util.APIResponse(ctx, err.Error(), http.StatusBadRequest, http.MethodPost, nil)
	} else {

		resultRegister, errRegister := h.service.RegisterService(&input)

		switch errRegister {

		case "REGISTER_CONFLICT_409":
			util.APIResponse(ctx, "Username already exist", http.StatusConflict, http.MethodPost, nil)
			return

		case "REGISTER_FAILED_403":
			util.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)
			return

		default:
			accessTokenData := map[string]interface{}{"id": resultRegister.ID, "username": resultRegister.Username}
			_, errToken := util.Sign(accessTokenData, util.GodotEnv("JWT_SECRET"), 60)

			if errToken != nil {
				defer logrus.Error(errToken.Error())
				util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
				return
			}

			util.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, input)
		}
	}
}
