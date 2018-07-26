package services

import (
	"database/sql"
	"landing-place/helpers"
	"landing-place/models"
)

func SignUp(user *models.User) error {
	db := helpers.Db

	row := db.QueryRow("INSERT INTO users(name, email, password) VALUES($1, $2, $3) returning id, password", &user.Name, &user.Email, &user.Password)

	err := row.Scan(&user.ID, &user.Password)

	return err
}

func SignIn(email string) *sql.Row {
	db := helpers.Db

	row := db.QueryRow("SELECT * FROM users WHERE email=$1;", email)

	return row
}
