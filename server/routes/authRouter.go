package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/nihalmenon/forum/controllers"
)

func AuthRoutes(incoming *gin.Engine) {
	// not protected routes (no middleware)
	users := incoming.Group("/users")
	{
		users.POST("/register", controller.Register())
		users.POST("/login", controller.Login())
	}
}
