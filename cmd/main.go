package main

import (
	"estudosGo/internal/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/users", controller.GetUsers)

	server.Run()
}
