package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RequestRouter(r *gin.Engine) {
	router := r.Group("/get")

	router.GET("param/:email/:password", func(ctx *gin.Context) {
		email := ctx.Param("email")
		password := ctx.Param("password")

		ctx.JSON(http.StatusOK, gin.H{
			"email":    email,
			"password": password,
		})
	})

	router.GET("query", func(ctx *gin.Context) {
		email := ctx.Query("email")
		password := ctx.Query("password")

		ctx.JSON(http.StatusOK, gin.H{
			"email":    email,
			"password": password,
		})
	})

	router.GET("query/binding", func(ctx *gin.Context) {
		email := ctx.Query("email")
		password := ctx.QueryArray("password")

		type Request struct {
			Email    string   `form:"email"`
			Password []string `form:"password"`
		}

		var res Request

		err := ctx.ShouldBindWith(&res, binding.Query)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
		}
		ctx.Header("email", "CUANBOT")
		ctx.JSON(http.StatusOK, gin.H{
			"email":    email,
			"password": password,
			"data":     res,
		})
	})
}
