package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishanmadhav/enmesh/controllers"
)

func ProfileRoutes(router *gin.RouterGroup) {
	profile := router.Group("/profile")

	profile.GET("/:profileID", controllers.GetProfileByID)
	profile.GET("/all", controllers.GetAllProfiles)
	profile.POST("/", controllers.CreateProfile)
	profile.DELETE("/:profileID", controllers.DeleteProfileByID)

}
