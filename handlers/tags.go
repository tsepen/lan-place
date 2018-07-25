package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"landing-place/helpers"
	"landing-place/models"
	"landing-place/services"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Tag struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func GetTags(w http.ResponseWriter, r *http.Request) {
	db := helpers.Db

	rows, err := db.Query("SELECT * FROM tags")
	if err != nil {
		http.Error(w, "Not found", 400)
		return
	}

	defer rows.Close()

	tags := make([]*Tag, 0)
	for rows.Next() {
		tag := new(Tag)
		err := rows.Scan(&tag.ID, &tag.Title)
		if err != nil {
			fmt.Println(err)
		}

		tags = append(tags, tag)
	}

	js, _ := json.MarshalIndent(tags, "", " ")

	w.Write(js)
}

func GetOneTag(w http.ResponseWriter, r *http.Request) {
	var tag Tag

	vars := mux.Vars(r)

	tag.ID = vars["id"]

	db := helpers.Db

	row := db.QueryRow("SELECT * FROM tags WHERE id=$1", tag.ID)

	err := row.Scan(&tag.ID, &tag.Title)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	js, _ := json.MarshalIndent(tag, "", " ")

	w.Write(js)

}

func CreateTag(w http.ResponseWriter, r *http.Request) {
	tag := models.Tag{
		Title: r.FormValue("title"),
	}

	err := services.CreateTag(&tag)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	js, err := json.MarshalIndent(tag, "", " ")

	w.Write(js)
}

func UpdateTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tag := Tag{
		ID:    vars["id"],
		Title: r.FormValue("title"),
	}

	db := helpers.Db

	result, err := db.Exec("UPDATE tags SET title = $1 WHERE id = $2", tag.Title, tag.ID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil || rowsAffected == 0 {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	js, _ := json.Marshal(tag)

	w.Write(js)
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	db := helpers.Db

	result, err := db.Exec("DELETE FROM tags WHERE id = $1", id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil || rowsAffected == 0 {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Tag with id %s deleted successfully\n", id)
}
