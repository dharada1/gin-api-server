package controllers

import (
	"github.com/yhirano55/gin-api-server/models"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (c User) Get(id int) interface{} {
	repo := models.NewUserRepository()
	user := repo.GetByID(id)
	return user
}
