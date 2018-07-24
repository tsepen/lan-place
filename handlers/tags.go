package handlers

import (
	"encoding/json"
	"fmt"
	"landing-place/helpers"
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

	err := db.QueryRow("SELECT * FROM tags WHERE id=$1", tag.ID).Scan(&tag.ID, &tag.Title)
	if err != nil {
		http.Error(w, "Not found", 400)
		return
	}
	js, _ := json.MarshalIndent(tag, "", " ")

	w.Write(js)

}

func CreateTag(w http.ResponseWriter, r *http.Request) {
	tag := Tag{
		Title: r.FormValue("title"),
	}

	db := helpers.Db

	err := db.QueryRow("INSERT INTO tags(title) VALUES($1) returning id", tag.Title).Scan(&tag.ID)
	if err != nil {
		http.Error(w, "This tag in database", 400)
		return
	}
	js, _ := json.MarshalIndent(tag, "", " ")

	w.Write(js)
}

func UpdateTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tag := Tag{
		ID:    vars["id"],
		Title: r.FormValue("title"),
	}

	db := helpers.Db

	db.QueryRow("UPDATE tags SET title = $1 WHERE id = $2", tag.Title, tag.ID).Scan(&tag.ID, &tag.Title)

	js, _ := json.Marshal(tag)

	w.Write(js)
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	db := helpers.Db

	db.QueryRow("DELETE FROM tags WHERE id = $1", id)

	fmt.Fprint(w, "Delete")
}
