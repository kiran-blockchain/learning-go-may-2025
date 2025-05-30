package dependency

import (
    "rest-api/controllers"
    "rest-api/repositories"
    "rest-api/services"
)

type Container struct {
    UserController *controllers.UserController
    OrderController *controllers.OrderController
}

func BuildContainer() *Container {
	// 1. Model ->Create Interface -> Repository ->Service  -> Controller -> Route

    // Wiring dependencies manually
    userRepo := repositories.NewUserRepository()
    userService := services.NewUserService(userRepo)
    userController := controllers.NewUserController(userService)

    orderRepo := repositories.NewOrderRepository()
    orderService := services.NewMatchingService(orderRepo)
    orderController := controllers.NewOrderController(orderService)


    return &Container{
        UserController: userController,
        OrderController: orderController,
    }
}
