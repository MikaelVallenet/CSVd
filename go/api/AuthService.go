package api

import (
	"github.com/Mikatech/CTF-AI/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginUser struct {
	Email    string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (env Env) Login(c *gin.Context) {
	var user LoginUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userDb model.User
	if userDb, err := env.Svc.GetUserByMail(user.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO add jwt
}
