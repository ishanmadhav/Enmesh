package controllers

import "github.com/gin-gonic/gin"

func GetTweet(c *gin.Context) {
	c.JSON(200, gin.H{
		"body": "Tweet content",
	})
}

func GetAllTweets(c *gin.Context){
	c.JSON(200, gin.H{
		"body": "List of all tweets",
	})
}

func DeleteAllTweets(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "All tweets deleted",
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