package controllers

import (
	"github.com/JackMaarek/cqrsPattern/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	cmd, err := services.CreateUser(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		c.Abort()
	}
	c.JSON(http.StatusCreated, cmd.Data)
	return
}
