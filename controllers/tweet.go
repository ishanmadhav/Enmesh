package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishanmadhav/enmesh/database"
	"github.com/ishanmadhav/enmesh/models"
)

type TweetBody struct {
	Body      string `json:"body"`
	ProfileID uint   `json:"profileID"`
}

func GetTweetByID(c *gin.Context) {
	fmt.Println("Get Tweet by ID: " + c.Param("tweetID"))
}

func CreateTweet(c *gin.Context) {
	fmt.Println("Create tweet")
	db := database.DBConn
	var tempTweet TweetBody
	c.Bind(&tempTweet)
	newTweet := models.Tweet{Body: tempTweet.Body, ProfileID: tempTweet.ProfileID}
	result := db.Create((&newTweet))
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Print(newTweet)
	c.JSON(http.StatusOK, gin.H{
		"body": newTweet,
	})
}

func GetAllTweets(c *gin.Context) {
	fmt.Println("Get All Tweets")
}
