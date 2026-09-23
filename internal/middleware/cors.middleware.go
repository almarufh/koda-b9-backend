package middleware

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {
	allowedOrigins := []string{
		"http://localhost:8080",
		"http://localhost:3000",
		"http://localhost:8081",
		"http://127.0.0.1:3000",
	}
	fmt.Println(c.GetHeader("Origin"))
	fmt.Println(http.MethodOptions)
	if slices.Contains(allowedOrigins, c.GetHeader("Origin")) {
		c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	}
	// c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Methods", "OPTIONS, PATCH")

	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}
