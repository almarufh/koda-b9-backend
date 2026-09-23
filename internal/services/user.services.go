package services

import (
	"fmt"

	"github.com/almarufh/koda-b9-backend/internal/dto"
)

type ServiceUser struct {
	users []dto.User
}

func NewUserService() *ServiceUser {
	return &ServiceUser{
		users: []dto.User{
			{
				Name:     "Alma'ruf Hidayat",
				Email:    "admin@mail.com",
				Password: "12345678",
			},
		},
	}
}

func (su *ServiceUser) AddUser(data dto.User) error {
	if len(data.Password) < 8 {
		return fmt.Errorf("length passwor minimum 8 character")
	}

	user, err := su.GetUserByEmail(data.Email)
	if err == nil {
		return fmt.Errorf("email alredy exist !")
	}

	su.users = append(su.users, user)
	return nil
}

func (su *ServiceUser) GetUserByEmail(e string) (dto.User, error) {
	for _, v := range su.users {
		if v.Email == e {
			return v, nil
		}
	}
	return dto.User{}, fmt.Errorf("%s not found !", e)
}

func (su *ServiceUser) LoginServices(data dto.Login) (dto.User, error) {
	user, err := su.GetUserByEmail(data.Email)
	if err != nil {
		return dto.User{}, fmt.Errorf("%s", err)
	}

	if data.Password != user.Password {
		return dto.User{}, fmt.Errorf("wrong password !")
	}
	return user, nil
}
