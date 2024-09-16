package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// default gin
	r := gin.Default()

	// Directorio Javascript & CSS
	r.Static("/static", "./static")

	// Directorio Templates
	r.LoadHTMLGlob("templates/*")

	// Function call routes
	setupRoutes(r)

	// Corre en puerto 8080
	r.Run()
}
