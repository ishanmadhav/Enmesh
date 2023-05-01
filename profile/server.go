package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishanmadhav/Enmesh/profile/database"
	"github.com/ishanmadhav/Enmesh/profile/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=212699 dbname=enmesh port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routes.ProfileRoutes(router.Group("api"))
	router.Run(":8081")
}
