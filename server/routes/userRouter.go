package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/nihalmenon/forum/controllers"
	"github.com/nihalmenon/forum/middleware"
)

func UserRoutes(incoming *gin.Engine) {

	incoming.Use(middleware.AuthMiddleware())
	users := incoming.Group("/users")
	{
		users.GET("/", controller.GetUsers())
		users.GET("/:user_id", controller.GetUser())
		// users.POST("/", controller.CreateUser())
		// users.PUT("/:id", controller.UpdateUser())
		// users.DELETE("/:id", controller.DeleteUser())
	}
}
