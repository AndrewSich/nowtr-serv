package main

import (
	"net/http"

	"nowtr/configs"
	"nowtr/users"

	"github.com/gin-gonic/gin"
)

func main() {
	serv := gin.Default()
	serv.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "success",
			"data":    "Welcome to Nowtr Main Server",
		})
	})

	// config
	configs.ConnectGoerli()

	routerApps := serv.Group("/")
	users.RouterUsers(routerApps.Group("/users"))

	// running on server
	serv.Run()
}
