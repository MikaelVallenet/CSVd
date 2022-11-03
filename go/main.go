package main

import (
	"github.com/Mikatech/CTF-AI/controller"
	"github.com/Mikatech/CTF-AI/database"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	db, err := database.Database()
	if err != nil {
		log.Fatal(err)
		return
	}
	db.DB()

	router := gin.Default()

	router.GET("/user/:id", controller.GetUser)
	router.GET("/users", controller.GetUsers)
	router.POST("/user", controller.CreateUser)
	router.DELETE("/user/:id", controller.DeleteUser)
	log.Fatal(router.Run(":10000"))
}
