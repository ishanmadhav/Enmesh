package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishanmadhav/enmesh/controllers"
)

func TweetRoutes(router *gin.RouterGroup) {
	tweet := router.Group("/tweet")
	tweet.GET("/:tweetID", controllers.GetTweetByID)
	tweet.POST("/", controllers.CreateTweet)
	tweet.GET("/all", controllers.GetAllTweets)

}
