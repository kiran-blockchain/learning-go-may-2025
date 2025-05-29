package routes

import (
	"rest-api/controllers"
	"github.com/gin-gonic/gin"

)

func BuildRoutes(r *gin.Engine) {
	r.POST("/create", controllers.CreateUser)
	//route
	r.GET("/get", controllers.GetUsers)
}
