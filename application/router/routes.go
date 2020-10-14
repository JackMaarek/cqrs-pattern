package router

import (
	"github.com/JackMaarek/cqrsPattern/application/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	apiV1 := router.Group("/v1")

	//>------------------ Users --------------------------
	users := apiV1.Group("/users")
	users.GET("/", controllers.GetUserList)
	users.GET("/:id", controllers.GetUser)
	users.POST("/", controllers.CreateUser)
	users.PUT("/:id", controllers.UpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)
	//<------------------ End Users --------------------------

	//>------------------ Orders --------------------------
	orders := apiV1.Group("/orders")
	orders.POST("/", controllers.CreateOrder)
	//<------------------ End Orders --------------------------

}
