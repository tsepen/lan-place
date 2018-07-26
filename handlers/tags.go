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
	_ "github.com/lib/pq"
)

func GetTags(w http.ResponseWriter, r *http.Request) {

	rows, err := services.GetTag()
	if err != nil {
		http.Error(w, "Not found", 400)
		return
	}

	defer rows.Close()

	tags := make([]*models.Tag, 0)
	for rows.Next() {
		tag := new(models.Tag)
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
	var tag models.Tag

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}

	tag.ID = id

	row := services.GetOneTag(tag.ID)

	err = row.Scan(&tag.ID, &tag.Title)
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

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}

	tag := models.Tag{
		ID:    id,
		Title: r.FormValue("title"),
	}

	rowsAffected, err := services.UpdateTag(&tag)
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

	rowsAffected, err := services.DeleteTag(id)
	if err != nil || rowsAffected == 0 {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Tag with id %s deleted successfully\n", id)
}
