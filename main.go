package main

import (
	"github.com/Morizz00/go-otp-sender/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	app := api.Config{Router: router}

	app.Routes()

	router.Run(":8000")
}
