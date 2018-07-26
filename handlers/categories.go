package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"landing-place/models"
	"landing-place/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Category struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func GetCategories(w http.ResponseWriter, r *http.Request) {

	rows, err := services.GetCategory()
	if err != nil {
		http.Error(w, "Not found", 400)
		return
	}

	defer rows.Close()

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

	row := services.GetOneCategory(category.ID)

	err := row.Scan(&category.ID, &category.Title)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	js, _ := json.MarshalIndent(category, "", " ")

	w.Write(js)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	category := models.Category{
		Title: r.FormValue("title"),
	}

	err := services.CreateCategory(&category)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	js, _ := json.MarshalIndent(category, "", " ")

	w.Write(js)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}

	category := models.Category{
		ID:    id,
		Title: r.FormValue("title"),
	}

	rowsAffected, err := services.UpdateCategory(&category)
	if err != nil || rowsAffected == 0 {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	js, _ := json.Marshal(category)

	w.Write(js)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	rowsAffected, err := services.DeleteCategory(id)
	if err != nil || rowsAffected == 0 {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Category with id %s deleted sucsessfuly\n", id)

}
