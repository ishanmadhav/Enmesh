package controllers

import "github.com/gin-gonic/gin"

func GetTweet(c *gin.Context) {
	c.JSON(200, gin.H{
		"body": "Tweet content",
	})
}

func PostTweet(c *gin.Context) {
	c.JSON(200, gin.H{
		"body": "Tweet content",
	})
}

func EditTweet(c *gin.Context){
	c.JSON(200, gin.H{
		"body": "Tweet content",
	})
}

func DeleteTweet(c *gin.Context){
	c.JSON(200, gin.H{
		"body": "Tweet content",
	})
}