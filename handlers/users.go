package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"landing-place/helpers"
	"landing-place/models"
	"landing-place/services"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	var u = new(models.User)

	row := services.SignIn(user.Email)

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
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
	user := models.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	hash, _ := helpers.HashPassword(user.Password)

	user.Password = hash

	err := services.SignUp(&user)
	if err != nil {
		http.Error(w, "This email register", 400)
		return
	}

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
