package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishanmadhav/enmesh/database"
	"github.com/ishanmadhav/enmesh/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=mysecretpassword dbname=enmesh port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	routes.TweetRoutes(api)
	routes.ProfileRoutes(api)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
