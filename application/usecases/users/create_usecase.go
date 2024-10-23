package users

import (
	"github.com/guilhermelinosp/crud1/application/services/cryptography"
	"github.com/guilhermelinosp/crud1/config/errs"
	"github.com/guilhermelinosp/crud1/config/logs"
	"github.com/guilhermelinosp/crud1/domain/dtos/requests"
)

func CreateTask(user *requests.UserRequest) *errs.Error {
	logs.Info("Init UseCases.Users.CreateTask")

	encryptedPassword, err := cryptography.EncryptPassword(user.Password)
	if err != nil {
		logs.Error("Failed to encrypt user password", err)
		return errs.NewBadRequest("Error occurred while encrypting the password")
	}

	user.Password = encryptedPassword

	// Code for creating user
	// userRepo := repositories.NewUserRepository()
	// userRepo.CreateUser(user)

	return nil
}
