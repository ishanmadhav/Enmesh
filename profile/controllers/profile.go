package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ishanmadhav/Enmesh/models"
	"github.com/ishanmadhav/Enmesh/profile/database"
)

func GetProfileById(c *gin.Context) {
}

func GetAllProfiles(c *gin.Context) {
	c.JSON(200, gin.H{
		"body": "Profile",
	})
}

func CreateProfile(c *gin.Context) {
	db:=database.DBConn
	profile:=&models.Profile{Name: "John", Age: 32, Email: "john@gmail.com", Username: "john", Tweets: []}
}

func DeleteProfileById(c *gin.Context) {

}
