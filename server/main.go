package main

import (
	"net/http"

	routes "github.com/nihalmenon/forum/routes"

	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	log.Println("Starting server...")

	router := gin.Default()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from the API",
		})
	})

	router.Run(":" + PORT)
}
