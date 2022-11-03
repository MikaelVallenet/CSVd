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
	
	router.GET("/users", controller.getUsers)
	log.Fatal(router.Run(":10000"))
}
