package router

import (
	"github.com/JackMaarek/cqrsPattern/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.POST("/createUser", controllers.CreateUser)
}
