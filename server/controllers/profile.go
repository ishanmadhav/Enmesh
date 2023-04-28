package controllers

import "github.com/gin-gonic/gin"

func GetProfileById(c *gin.Context){
	c.JSON(200, gin.H{
		"body": "Profile",
	})
}

func GetAllProfiles(c *gin.Context){
	c.JSON(200, gin.H{
		"body": "List of all profiles",
	})
}

func CreateProfile(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "New profile created",
	})
}

func DeleteProfileById(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Profile deleted",
	})
}