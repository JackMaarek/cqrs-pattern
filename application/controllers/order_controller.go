package controllers

import (
	"github.com/JackMaarek/cqrsPattern/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	cmd, err := services.CreateOrder(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusCreated, cmd)
	return
}
