package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", hello)
	engine.Run(":7777")
	// 321111
}

// Handler
func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, World")
}
