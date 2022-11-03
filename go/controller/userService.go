package controller

import (
	"github.com/Mikatech/CTF-AI/database"
	"github.com/Mikatech/CTF-AI/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type NewUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetUsers(c *gin.Context) {

	var users []model.User

	db, err := database.Database()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {

	var user model.User

	db, err := database.Database()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {

	var user NewUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := model.User{Name: user.Name, Password: user.Password}

	db, err := database.Database()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func DeleteUser(c *gin.Context) {

	var user model.User

	db, err := database.Database()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
