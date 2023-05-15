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
	db := database.DBConn
	var profile models.Profile
	result := db.Model(models.Profile{}).Preload("Tweets").Find(&profile, c.Param("profileID"))
	if result.Error != nil {
		fmt.Print(result.Error)
	}
	c.JSON(http.StatusOK, gin.H{
		"body": profile,
	})
}

func GetAllProfiles(c *gin.Context) {
	fmt.Println("Get All Profiles")
	db := database.DBConn
	var profiles []models.Profile
	result := db.Find(&profiles)
	if result.Error != nil {
		panic("Fetching all profiles did not work")
	}
	fmt.Print(profiles)
	c.JSON(http.StatusOK, gin.H{
		"body": profiles,
	})
}

func CreateProfile(c *gin.Context) {
	db := database.DBConn
	fmt.Println("Create Profile Route hit")
	var tempProfileBody ProfileBody
	c.Bind(&tempProfileBody)
	newProfile := models.Profile{Name: tempProfileBody.Name, Age: tempProfileBody.Age, Username: tempProfileBody.Username}
	fmt.Print(newProfile)
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
