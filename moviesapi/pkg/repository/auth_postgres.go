package repository

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	moviesapi "github.com/LiaNestazi/moviesApi"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user moviesapi.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password, role) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Username, user.Password, "user")
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (moviesapi.User, error) {
	var user moviesapi.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}