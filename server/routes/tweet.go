package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishanmadhav/Enmesh/server/controllers"
)

func TweetRoutes(router *gin.RouterGroup) {
	tweet := router.Group("/tweet")
	{
		tweet.GET("/", controllers.GetTweet)
		tweet.GET("/all", controllers.GetAllTweets)
		tweet.POST("/", controllers.PostTweet)
		tweet.PUT("/", controllers.EditTweet)
		tweet.DELETE("/", controllers.DeleteTweet)
		tweet.DELETE("/", controllers.DeleteAllTweets)
	}

}
