package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/", hello)
	engine.Run(":7777")
	//
}

// Handler
func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, World!")
	//
}
