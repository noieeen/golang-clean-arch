package usecases

import (
	"auth-example/modules/entities"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type usersUse struct {
	UserRepo entities.UsersRepository
}

// Constructor
func NewUsersUsecase(userRepo entities.UsersRepository) entities.UsersUsecase {
	return &usersUse{
		UserRepo: userRepo,
	}
}

func (u *usersUse) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		fmt.Println((err.Error()))
		return nil, err
	}

	req.Password = string(hashed)

	// Send req next to repository
	user, err := u.UserRepo.Register(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}
