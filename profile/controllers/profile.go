package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ishanmadhav/Enmesh/profile/database"
	"github.com/ishanmadhav/Enmesh/profile/models"
)

func GetProfileById(c *gin.Context) {
}

func GetAllProfiles(c *gin.Context) {
	c.JSON(200, gin.H{
		"body": "Profile",
	})
}

func CreateProfile(c *gin.Context) {
	db := database.DBConn
	var emptyTweetsArr []models.Tweet
	profile := models.Profile{Name: "John", Email: "john@gmail.com", Username: "john2", Age: 32, TweetsList: emptyTweetsArr}
	result := db.Create(&profile)
	fmt.Print(result)
	c.JSON(200, gin.H{
		"body": "Profile",
	})
}

func DeleteProfileById(c *gin.Context) {

}
