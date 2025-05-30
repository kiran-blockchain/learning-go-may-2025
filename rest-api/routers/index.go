package routers

import (
    "github.com/gin-gonic/gin"
    "rest-api/controllers"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {
    r := gin.Default()

    // Users endpoint
    r.GET("/users", userController.GetAllUsers)

    return r
}
