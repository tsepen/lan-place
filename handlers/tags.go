package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get tags")
}

func GetOneTag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one tag")
}

func CreateTag(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	js, _ := json.MarshalIndent(user, "", " ")

	w.Write(js)
}

func UpdateTag(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	js, _ := json.MarshalIndent(user, "", " ")

	w.Write(js)
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	js, _ := json.MarshalIndent(user, "", " ")

	w.Write(js)
}
