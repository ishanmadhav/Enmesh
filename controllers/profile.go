package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishanmadhav/enmesh/database"
	"github.com/ishanmadhav/enmesh/models"
)

type ProfileBody struct {
	Name     string `json:"name"`
	Age      uint   `json:"age"`
	Username string `json:"username"`
}

func GetProfileByID(c *gin.Context) {
	fmt.Println("Get Profile By Id")
}

func GetAllProfiles(c *gin.Context) {
	fmt.Println("Get All Profiles")
}

func CreateProfile(c *gin.Context) {
	db := database.DBConn
	fmt.Println("Create Profile Route hit")
	var tempProfileBody ProfileBody
	c.Bind(&tempProfileBody)
	newProfile := models.Profile{Name: tempProfileBody.Name, Age: tempProfileBody.Age, Username: tempProfileBody.Username}
	result := db.Create(&newProfile)
	if result.Error != nil {
		fmt.Print(result.Error)
	} else {
		fmt.Print(result)
	}
	c.JSON(http.StatusOK, gin.H{
		"body": newProfile,
	})

}

func DeleteProfileByID(c *gin.Context) {
	fmt.Println("Deleted Profile")
}
