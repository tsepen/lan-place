package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get categories")
}

func GetOneCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one cat")
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	js, _ := json.MarshalIndent(user, "", " ")

	w.Write(js)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	js, _ := json.MarshalIndent(user, "", " ")

	w.Write(js)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	js, _ := json.MarshalIndent(user, "", " ")

	w.Write(js)
}
