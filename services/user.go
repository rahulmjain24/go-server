package services

import (
	"github.com/rahulmjain24/go-server/models"
	"github.com/rahulmjain24/go-server/repositories"
)

func RegisterUser(name, email string) error {
	return repositories.CreateUser(name, email)
}

func FindUserByEmail(email string) (*models.User, error) {
	return repositories.GetUserByEmail(email)
}
