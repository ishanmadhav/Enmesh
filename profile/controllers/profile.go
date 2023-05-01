package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetProfileById(c *gin.Context) {

}

func GetAllProfiles(c *gin.Context) {
	c.JSON(200, gin.H{
		"body": "Profile",
	})
}

func CreateProfile(c *gin.Context) {

}

func DeleteProfileById(c *gin.Context) {

}
