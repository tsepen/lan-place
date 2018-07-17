package handlers

import (
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
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	fmt.Println(user)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	db := helpers.Db
	var id int
	err := db.QueryRow("INSERT INTO users(name, email, password) VALUES($1, $2, $3) returning id", user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}

	user.ID = id
	fmt.Println(user)
}
