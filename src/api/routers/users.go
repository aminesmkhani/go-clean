package routers

import (
	"github.com/aminesmkhani/go-clean/api/handlers"
	"github.com/aminesmkhani/go-clean/api/middlewares"
	"github.com/aminesmkhani/go-clean/config"
	"github.com/gin-gonic/gin"
)


func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)
	router.POST("/send-otp", middlewares.OtpLimiter(cfg), func(c *gin.Context) {
		h.SendOtp(c)
	})
}