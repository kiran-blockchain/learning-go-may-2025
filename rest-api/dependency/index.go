package dependency

import (
    "rest-api/controllers"
    "rest-api/repositories"
    "rest-api/services"
)

type Container struct {
    UserController *controllers.UserController
}

func BuildContainer() *Container {
	// 1. Model ->Create Interface -> Repository ->Service  -> Controller -> Route

    // Wiring dependencies manually
    userRepo := repositories.NewUserRepository()
    userService := services.NewUserService(userRepo)
    userController := controllers.NewUserController(userService)

    return &Container{
        UserController: userController,
    }
}
