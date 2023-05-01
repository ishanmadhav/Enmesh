package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishanmadhav/Enmesh/profile/controllers"
)

func ProfileRoutes(router *gin.RouterGroup) {
	profile := router.Group("/profile")
	profile.GET("/:profileId", controllers.GetProfileById)
	profile.GET("/all", controllers.GetAllProfiles)
	profile.POST("/", controllers.CreateProfile)
	profile.DELETE("/", controllers.DeleteProfileById)
}
