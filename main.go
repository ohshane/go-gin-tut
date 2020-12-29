package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.Default()

func main() {
	// router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	router.Run()
}
