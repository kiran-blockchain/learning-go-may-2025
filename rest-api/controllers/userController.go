package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "mypost",
		"abc":     "xyz",
	})

}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "mypost",
		"abc":     "xyz",
	})

}

func DeleteUser() {

}

func UpdateUser() {

}
