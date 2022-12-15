package api

import (
	"github.com/Mikatech/CTF-AI/database"
	"github.com/Mikatech/CTF-AI/model"
	"github.com/Mikatech/CTF-AI/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Env struct {
	Svc *database.UserService
}

func (env Env) GetUsers(c *gin.Context) {
	users, err := env.Svc.GetUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (env Env) GetUser(c *gin.Context) {
	user, err := env.Svc.GetUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (env Env) CreateUser(c *gin.Context) {
	var user NewUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = tools.EncryptPassword(user.Password)
	newUser := model.User{Email: user.Email, Password: user.Password}
	if err := env.Svc.CreateUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func (env Env) DeleteUser(c *gin.Context) {
	user, err := env.Svc.GetUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	err = env.Svc.DeleteUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
