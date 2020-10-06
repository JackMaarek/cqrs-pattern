package router

import (
	"github.com/JackMaarek/cqrsPattern/application/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	apiV1 := router.Group("/v1")

	users := apiV1.Group("/users")
	users.POST("/", controllers.CreateUser)
	users.GET("/", controllers.GetUserList)
	users.PUT("/:id", controllers.UpdateUser)
}
