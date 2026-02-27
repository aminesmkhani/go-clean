package api

import (
	"fmt"

	"github.com/aminesmkhani/go-clean/api/middlewares"
	"github.com/aminesmkhani/go-clean/api/routers"
	validation "github.com/aminesmkhani/go-clean/api/validations"
	"github.com/aminesmkhani/go-clean/config"
	"github.com/aminesmkhani/go-clean/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	// Register custom Validator!
	RegisterValidators()

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(gin.Logger(), gin.Recovery(),middlewares.LimitByRequest())
	r.Use(middlewares.Cors(cfg))
	r.Use(middlewares.DefaultStructuredLogger(cfg))

	// Register Routes
	RegisterRoutes(r,cfg)
	// Register Swagger
	RegisterSwagger(r,cfg)

	r.Run(fmt.Sprintf(":%s",cfg.Server.InternalPort))
}



func RegisterRoutes(r *gin.Engine,cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		users := v1.Group("/users")

		routers.Health(health)
		routers.TestRouter(test_router)
		routers.User(users,cfg)
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

func RegisterSwagger(r *gin.Engine,cfg *config.Config){
	docs.SwaggerInfo.Title = " Golang Clean Web Api"
	docs.SwaggerInfo.Description = " Golang web api for car sell"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s",cfg.Server.InternalPort)
	docs.SwaggerInfo.Schemes= []string{"http"}
	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
}