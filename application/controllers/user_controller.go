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
	c.JSON(http.StatusOK, ul)
	return

}

func CreateUser(c *gin.Context) {
	cmd, err := services.CreateUser(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		c.Abort()
	}
	c.JSON(http.StatusCreated, cmd)
	return
}