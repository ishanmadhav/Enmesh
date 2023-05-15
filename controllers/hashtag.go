package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishanmadhav/enmesh/database"
	"github.com/ishanmadhav/enmesh/models"
)

type HashtagBody struct {
	Name string `json:"name"`
}

func GetHashtagByTag(c *gin.Context) {

}

func CreateHashtag(c *gin.Context) {
	fmt.Println("Creating hashtag")
	db := database.DBConn
	var tempHashtag HashtagBody
	c.Bind(&tempHashtag)
	hashtag := models.Hashtag{Name: tempHashtag.Name}
	db.Create(&hashtag)
	c.JSON(http.StatusCreated, gin.H{
		"body": hashtag,
	})
}

func GetAllHashtags(c *gin.Context) {

}

func DeleteHashtagByTag(c *gin.Context) {

}
