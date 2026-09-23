package routers

import (
	"github.com/almarufh/koda-b9-backend/internal/controllers"
	"github.com/almarufh/koda-b9-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func initAuthRouters(r *gin.Engine) {
	us := services.NewUserService()
	uc := controllers.NewUserController(us)

	authRouters := r.Group("/auth")
	authRouters.POST("/login", uc.Login)
	authRouters.POST("/register", uc.Register)
}
