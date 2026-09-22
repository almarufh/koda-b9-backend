package main

import (
	"github.com/almarufh/koda-b9-backend/internal/routers"
	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  routers.Routers(r)
  r.Run(":8080")
}