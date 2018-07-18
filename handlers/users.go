package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"landing-place/helpers"

	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	user := User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	db := helpers.Db

	var u = new(User)

	row := db.QueryRow("SELECT * FROM users WHERE email=$1;", user.Email)
	err := row.Scan(&u.ID, &u.Email, &u.Name, &u.Password)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}

	match := helpers.CheckPasswordHash(user.Password, u.Password)

	if !match {
		http.Error(w, "Invalid login or password", 402)
	}
	if match {
		js, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

}

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	hash, _ := helpers.HashPassword(user.Password) // ignore error for the sake of simplicity

	user.Password = hash

	db := helpers.Db
	var id int
	var password string
	err := db.QueryRow("INSERT INTO users(name, email, password) VALUES($1, $2, $3) returning id, password", user.Name, user.Email, user.Password).Scan(&id, &password)
	if err != nil {
		fmt.Println(err)
	}

	user.ID = id
	user.Password = password

	js, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
