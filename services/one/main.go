package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jobrunner/sampletogo/services/one/server"
)

func main() {
	router := gin.Default()
	router.GET("/users/:userId", server.GetUserByID)
	router.Run(":8010")
}
