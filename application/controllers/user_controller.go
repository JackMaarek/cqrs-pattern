package controllers

import (
	"github.com/JackMaarek/cqrsPattern/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(c *gin.Context) {
	ul, err := services.GetUserList(c)
	if err != nil {
		return
	}
	if ul == nil {
		return
	}
	c.JSON(http.StatusOK, ul)
	return

}

func GetUser(c *gin.Context) {
	u, err := services.GetUser(c)
	if err != nil {
		return
	}
	if u == nil {
		return
	}
	c.JSON(http.StatusOK, u)
	return

}

func CreateUser(c *gin.Context) {
	cmd, err := services.CreateUser(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusCreated, cmd)
	return
}

func UpdateUser(c *gin.Context) {
	cmd, err := services.UpdateUser(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusCreated, cmd)
	return
}

func DeleteUser(c *gin.Context) {
	err := services.RemoveUser(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, "User deleted")
	return
}
