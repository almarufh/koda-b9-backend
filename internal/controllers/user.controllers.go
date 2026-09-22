package controllers

import (
	"fmt"
	"net/http"

	"github.com/almarufh/koda-b9-backend/internal/dto"
	"github.com/almarufh/koda-b9-backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type User struct {
	UserService *services.ServiceUser
}

func NewUserController (us *services.ServiceUser) *User {
	return  &User{
		UserService: us,
	}
}

func (u *User) Register ( c *gin.Context) {
	var data dto.User
	e := c.ShouldBindWith(&data, binding.JSON)

	if e != nil {
		c.JSON(http.StatusInternalServerError, dto.Response {
			Success: false,
			Message: e.Error(),
			Data: nil,
		})
		return 
	}

	if err := u.UserService.AddUser(data); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response {
			Success: false,
			Message: fmt.Sprintln(err),
			Data: nil,
		})
	}

	c.JSON(http.StatusOK, dto.Response {
		Success: true,
		Message: fmt.Sprintf("register success : %s", data.Name),
		Data: data,
	})
  }

  func (u *User) Login ( c *gin.Context) {
	var data dto.Login
	
	if e := c.ShouldBindWith(&data, binding.JSON); e != nil {
		c.JSON(http.StatusInternalServerError, dto.Response {
			Success: false,
			Message: e.Error(),
			Data: nil,
		})
		return
	}

	user, err := u.UserService.LoginServices(data); 
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response {
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response {
		Success: true,
		Message: "login success",
		Data: user,
	})
  }