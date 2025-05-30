package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "rest-api/services"
)

type UserController struct {
    service services.UserService
}

func NewUserController(s services.UserService) *UserController {
    return &UserController{service: s}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
    users := c.service.GetAllUsers()
    ctx.JSON(http.StatusOK, users)
}
