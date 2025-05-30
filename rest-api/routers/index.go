package routers

import (
    "github.com/gin-gonic/gin"
    "rest-api/controllers"
)

func SetupRouter(userController *controllers.UserController,orderController *controllers.OrderController) *gin.Engine {
    r := gin.Default()

    // Users endpoint
    r.GET("/users", userController.GetAllUsers)

    r.POST("/orders", orderController.PlaceOrder)

    return r
}
