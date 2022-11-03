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

func getUsers(c *gin.Context) {

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
