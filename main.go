package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type User struct {
	Name string
	Email string
	Password string
}

type Login struct {
	Email string
	Password string
}

type Response struct {
  Success bool
  Message string
  Data any
}

var Users = []User{
	{
		Name: "Alma'ruf Hidayat",
		Email: "admin@mail.com",
		Password: "12345678",
	},
}

func main() {

  r := gin.Default()

  r.POST("/login", func( c *gin.Context) {
	var data Login
	e := c.ShouldBindWith(&data, binding.JSON)

	if e != nil {
		c.JSON(http.StatusInternalServerError, Response {
			Success: false,
			Message: e.Error(),
			Data: nil,
		})
		return
	}

	for _,v := range Users {
		if data.Email == v.Email && data.Password == v.Password {
			c.JSON(http.StatusOK, Response {
				Success: true,
				Message: fmt.Sprintf("login success : %s", v.Name),
				Data: data,
			})
			return
		}
	}

	c.JSON(http.StatusOK, Response {
		Success: false,
		Message: fmt.Sprintln("email atau password salah"),
		Data: nil,
	})
  })

  r.POST("/register", func( c *gin.Context) {
	var data User
	e := c.ShouldBindWith(&data, binding.JSON)

	if e != nil {
		c.JSON(http.StatusInternalServerError, Response {
			Success: false,
			Message: e.Error(),
			Data: nil,
		})
		return
	}

	if len(data.Password) < 8 {
		c.JSON(http.StatusInternalServerError, Response {
			Success: false,
			Message: fmt.Sprintln("length passwor minimum 8 character"),
			Data: nil,
		})
		return
	}

	for _,v := range Users {
		if data.Email == v.Email {
			c.JSON(http.StatusBadRequest, Response {
				Success: false,
				Message: fmt.Sprintf("Email alredy exist : %s", v.Email),
				Data: nil,
			})
			return
		}
	}

	Users = append(Users, data)
	c.JSON(http.StatusOK, Response {
		Success: true,
		Message: fmt.Sprintf("register success : %s", data.Name),
		Data: data,
	})
  })

  r.Run(":8080")
}