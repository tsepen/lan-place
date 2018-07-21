package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetSites(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get")
}

func GetOneSite(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getOne")
}

func CreateSite(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	js, _ := json.MarshalIndent(user, "", " ")

	w.Write(js)
}

func UpdateSite(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	js, _ := json.MarshalIndent(user, "", " ")

	w.Write(js)
}

func DeleteSite(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	js, _ := json.MarshalIndent(user, "", " ")

	w.Write(js)
}

func LikeSite(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	js, _ := json.MarshalIndent(user, "", " ")

	w.Write(js)
}
