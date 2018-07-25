package handlers

import (
	"context"
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
	var u = new(User)
	db := helpers.Db

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
		helpers.CreateSession(w, r, u.ID)
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

	hash, _ := helpers.HashPassword(user.Password)

	user.Password = hash

	db := helpers.Db
	var id int
	var password string
	row := db.QueryRow("INSERT INTO users(name, email, password) VALUES($1, $2, $3) returning id, password", user.Name, user.Email, user.Password)
	err := row.Scan(&id, &password)
	if err != nil {
		http.Error(w, "This email register", 400)
		return
	}

	user.ID = id
	user.Password = password

	js, _ := json.MarshalIndent(user, "", " ")

	w.Write(js)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	helpers.DestroySession(w, r)
}

func Auth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user := helpers.GetSession(w, r)
		if user.ID != 0 {
			ctx := context.WithValue(r.Context(), "user", user)
			f.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	}
}
