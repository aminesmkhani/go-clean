package api

import (
	"fmt"

	"github.com/aminesmkhani/go-clean/api/middlewares"
	"github.com/aminesmkhani/go-clean/api/routers"
	validation "github.com/aminesmkhani/go-clean/api/validations"
	"github.com/aminesmkhani/go-clean/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	// Register custom Validator!
	RegisterValidators()

	r.Use(gin.Logger(), gin.Recovery(),middlewares.LimitByRequest())
	r.Use(middlewares.Cors(cfg))

	// Register Routes
	RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%s",cfg.Server.InternalPort))
}



func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")

		routers.Health(health)
		routers.TestRouter(test_router)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}
}
