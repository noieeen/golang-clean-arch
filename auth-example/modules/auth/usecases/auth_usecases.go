package usecases

import (
	"auth-example/configs"
	"auth-example/modules/entities"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type authUse struct {
	AuthRepo  entities.AuthRepository
	UsersRepo entities.UsersRepository
}

func NewAuthUsecase(authRepo entities.AuthRepository, usersRepo entities.UsersRepository) entities.AuthUsecase {
	return &authUse{
		AuthRepo:  authRepo,
		UsersRepo: usersRepo,
	}
}

func (u *authUse) Login(cfg *configs.Configs, req *entities.UsersCredentials) (*entities.UsersLoginRes, error) {
	user, err := u.UsersRepo.FindOneUser(req.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("error, password is invalid")

	}

	token, err := u.AuthRepo.SignUsersAccessToken(user)
	if err != nil {
		return nil, err
	}

	res := &entities.UsersLoginRes{
		AccessToken: token,
	}

	return res, nil
}
