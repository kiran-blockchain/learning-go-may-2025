package main

import (
    "rest-api/dependency"
    "rest-api/routers"
)

func main() {
    // Build DI container
    container := dependency.BuildContainer()

    // Setup router
    r := routers.SetupRouter(container.UserController,container.OrderController)


    // Start server
    r.Run(":4000")
}
