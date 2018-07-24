package handlers

import (
	"encoding/json"
	"fmt"
	"landing-place/helpers"
	"net/http"

	"github.com/gorilla/mux"
)

type Category struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	db := helpers.Db

	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		fmt.Println(err)
	}
	categories := make([]*Category, 0)

	for rows.Next() {
		category := new(Category)

		err := rows.Scan(&category.ID, &category.Title)
		if err != nil {
			fmt.Println(err)
		}

		categories = append(categories, category)
	}

	js, _ := json.MarshalIndent(categories, "", " ")

	w.Write(js)
}

func GetOneCategory(w http.ResponseWriter, r *http.Request) {
	category := new(Category)

	vars := mux.Vars(r)

	category.ID = vars["id"]

	db := helpers.Db

	err := db.QueryRow("SELECT * FROM categories WHERE id=$1", category.ID).Scan(&category.ID, &category.Title)
	if err != nil {
		http.Error(w, "Not found", 400)
		return
	}
	js, _ := json.MarshalIndent(category, "", " ")

	w.Write(js)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	category := Category{
		Title: r.FormValue("title"),
	}

	db := helpers.Db

	err := db.QueryRow("INSERT INTO categories(title) VALUES($1) returning id", category.Title).Scan(&category.ID)
	if err != nil {
		http.Error(w, "This category in database", 400)
		return
	}
	js, _ := json.MarshalIndent(category, "", " ")

	w.Write(js)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	category := Category{
		ID:    vars["id"],
		Title: r.FormValue("title"),
	}

	db := helpers.Db

	db.QueryRow("UPDATE categories SET title = $1 WHERE id = $2", category.Title, category.ID).Scan(&category.ID, &category.Title)

	js, _ := json.Marshal(category)

	w.Write(js)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	db := helpers.Db

	db.QueryRow("DELETE FROM categories WHERE id = $1", id)

	fmt.Fprint(w, "Delete category")

}
