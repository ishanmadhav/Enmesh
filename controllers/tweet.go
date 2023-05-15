package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Tweet struct {
	Name string `json:"name"`
	Roll string `json:"roll"`
}

func GetTweetByID(c *gin.Context) {
	fmt.Println("Get Tweet by ID: " + c.Param("tweetID"))
}

func CreateTweet(c *gin.Context) {
	fmt.Println("Create Tweet")
	var tweet Tweet
	c.Bind(&tweet)
	fmt.Print(tweet)
}

func GetAllTweets(c *gin.Context) {
	fmt.Println("Get All Tweets")
}
