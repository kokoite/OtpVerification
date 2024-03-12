package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pranjal/golang-otp-handler/api"
)

func main() {
	router := gin.Default()

	app := api.Config{Router: router}
	app.Routes()
	router.Run(":8080")
}
