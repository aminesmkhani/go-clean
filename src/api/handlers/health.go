package handlers

import (
	"net/http"

	"github.com/aminesmkhani/go-clean/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{

}

func NewHealthHandler() *HealthHandler{
	return &HealthHandler{}
}

func (h *HealthHandler)Health(ctx *gin.Context){
	ctx.JSON(http.StatusOK,helper.GenerateBaseResponse(
		"Working",
		true,
		0,
	))
			return
}