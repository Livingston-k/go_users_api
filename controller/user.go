package controller

import (
	"go-rest-api/config"
	"go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(user)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	config.DB.Where("id = ?", id).Delete(user)
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	config.DB.Where("id = ?", id).First(user)
	c.BindJSON(&user)
	config.DB.Save(user)
	c.JSON(200, user)
}
