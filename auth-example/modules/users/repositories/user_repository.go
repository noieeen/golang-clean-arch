package repositories

import (
	"auth-example/modules/entities"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	Db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) entities.UsersRepository {
	return &userRepo{
		Db: db,
	}
}

func (r *userRepo) FindOneUser(username string) (*entities.UsersPassport, error) {
	query := `SELECT
	"id",
	"username",
	"password"
	FROM "users"
	WHERE "username" = $1`

	res := new(entities.UsersPassport)
	if err := r.Db.Get(res, query, username); err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("error, user not found")
	}
	return res, nil

}

func (r *userRepo) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	query := `
	INSERT INTO "users"(
	"username",
	"password"
	)VALUES($1, $2)
	RETURNING "id", "username";
	`

	// Initial a user obj
	user := new(entities.UsersRegisterRes)

	// Query part
	rows, err := r.Db.Queryx(query, req.Username, req.Password)
	if (err) != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		if err := rows.StructScan(user); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

	}
	defer rows.Close()
	return user, nil
}
