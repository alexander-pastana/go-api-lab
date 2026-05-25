package repository

import (
	"database/sql"

	"github.com/alexander-pastana/go-api-lab/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository (connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) CreateUser(user model.User) error {
	query, err := ur.connection.Prepare("INSERT INTO users (username, password) VALUES ($1, $2)")
	if err != nil {
		return err
	}

	defer query.Close()

	_, err = query.Exec(user.Name, user.Password)
		if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserByName(username string) (*model.User, error) {
	query, err := ur.connection.Prepare("SELECT id, username, password from users WHERE username = $1")
		if err != nil {
		return nil, err
	}

	defer query.Close()

	var user model.User
	err = query.QueryRow(username).Scan(
		&user.ID,
		&user.Name,
		&user.Password,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, err
}