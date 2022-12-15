package api

import "github.com/gin-gonic/gin"

func ApplyRoutes(env Env, router *gin.Engine) {
	router.GET("/user/:id", env.GetUser)
	router.GET("/users", env.GetUsers)
	router.POST("/register", env.CreateUser)
	router.DELETE("/user/:id", env.DeleteUser)
}
