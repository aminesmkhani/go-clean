package handlers

import (
	"net/http"

	"github.com/aminesmkhani/go-clean/api/dto"
	"github.com/aminesmkhani/go-clean/api/helper"
	"github.com/aminesmkhani/go-clean/config"
	"github.com/aminesmkhani/go-clean/services"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	services *services.UserService
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	service := services.NewUserService(cfg)
	return &UsersHandler{services: service}
}

// SendOtp godoc
// @Summary Sent otp to user
// @Description Send Secure otp to user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/send-otp [post]
func (h *UsersHandler) SendOtp(c *gin.Context) error {
	req := new(dto.GetOtpRequest)
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return nil
	}
	err = h.services.SendOtp(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return nil
	}
	// TODO:call internal SMS Service
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
	return nil
}
