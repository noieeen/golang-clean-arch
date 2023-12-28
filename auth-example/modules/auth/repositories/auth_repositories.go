package repositories

import (
	"auth-example/modules/entities"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type authRepo struct {
	Db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) entities.AuthRepository {
	return &authRepo{
		Db: db,
	}
}

func (r *authRepo) SignUsersAccessToken(req *entities.UsersPassport) (string, error) {
	claims := entities.UserClaims{
		Id:       req.Id,
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "access_token",
			Subject:   uuid.NewString(),
			Audience:  []string{"users"},
		},
	}

	mySigningKey := os.Getenv("JWT_SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return ss, nil
}
