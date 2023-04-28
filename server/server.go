package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishanmadhav/Enmesh/server/routes"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routes.TweetRoutes(r.Group("api"))

	r.Run()
}
