package controllers

import (
	"github.com/gin-gonic/gin"
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
	db := database.DBConn

}

func DeleteProfileById(c *gin.Context) {

}
