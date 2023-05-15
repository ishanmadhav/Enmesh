package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishanmadhav/enmesh/controllers"
)

func HashtagRoutes(router *gin.RouterGroup) {
	profile := router.Group("/hashtag")

	profile.GET("/:tag", controllers.GetHashtagByTag)
	profile.GET("/all", controllers.GetAllHashtags)
	profile.POST("/", controllers.CreateHashtag)
	profile.DELETE("/:tag", controllers.DeleteHashtagByTag)

}
