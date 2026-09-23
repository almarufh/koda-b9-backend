package routers

import (
	"github.com/almarufh/koda-b9-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	r.Use(middleware.Cors)
	initAuthRouters(r)
	RequestRouter(r)
}
