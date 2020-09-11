package services

import (
	"fmt"
	"github.com/JackMaarek/cqrsPattern/bus"
	"github.com/JackMaarek/cqrsPattern/structs"
	"github.com/JackMaarek/cqrsPattern/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func CreateUser(c *gin.Context) (*structs.Command, error) {
	command := structs.Command{}
	header := structs.Header{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
	}
	if header.ContentType == "application/json" {
		c.JSON(http.StatusBadRequest, "Wrong content-type header")
		c.Abort()
	}
	userForm := structs.UserForm{}
	if err := c.ShouldBindJSON(&userForm); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		c.Abort()
	}
	err := validators.V.Struct(userForm)

	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e)
	}

	command.Data = userForm
	command.Type = "User Creation Command"
	var cmd *structs.Command
	cmd, err = bus.HandleRequest(&command)
	if err != nil {
		return nil, err
	}
	println(cmd)
	return cmd, nil
	//ch <- command
}
