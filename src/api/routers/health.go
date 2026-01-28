package routers

import (
	"github.com/aminesmkhani/go-clean/api/handlers"
	"github.com/gin-gonic/gin"
)



func Health(r *gin.RouterGroup){
	handler := handlers.NewHealthHandler()
	
	r.GET("/",handler.Health)
	r.POST("/",handler.HealthPost)
	r.POST("/:id",handler.HealthPostById)

}