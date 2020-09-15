package router

import (
	"github.com/JackMaarek/cqrsPattern/application/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.POST("/users", controllers.CreateUser)
	router.GET("users", controllers.GetUserList)
}
