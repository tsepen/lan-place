package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"landing-place/helpers"
)

type Site struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Created     string `json:"createdAt"`
	Rating      int    `json:"rating"`
	Views       int    `json:"views"`
}

func GetSites(w http.ResponseWriter, r *http.Request) {
	db := helpers.Db

	rows, err := db.Query("SELECT * FROM sites")
	if err != nil {
		fmt.Println(err)
	}

	sites := make([]*Site, 0)

	for rows.Next() {
		site := new(Site)

		err := rows.Scan(&site.ID, &site.Title, &site.Description, &site.URL, &site.Created, &site.Rating, &site.Views)
		if err != nil {
			fmt.Println(err)
		}

		sites = append(sites, site)
	}

	js, _ := json.MarshalIndent(sites, "", " ")

	w.Write(js)
}

func GetOneSite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	site := new(Site)

	db := helpers.Db

	err := db.QueryRow("UPDATE sites SET views=views+1 WHERE id=$1 returning *", id).Scan(&site.ID, &site.Title, &site.Description, &site.URL, &site.Created, &site.Rating, &site.Views)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Site not found", 404)
		return
	}

	js, _ := json.MarshalIndent(site, "", " ")

	w.Write(js)
}

func CreateSite(w http.ResponseWriter, r *http.Request) {
	site := Site{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		URL:         r.FormValue("url"),
	}

	db := helpers.Db

	err := db.QueryRow("INSERT INTO sites(title, description, url) VALUES($1, $2, $3) returning id, created, rating, views", site.Title, site.Description, site.URL).Scan(&site.ID, &site.Created, &site.Rating, &site.Views)
	if err != nil {
		http.Error(w, "This site in database", 400)
		return
	}
	js, _ := json.MarshalIndent(site, "", " ")

	w.Write(js)
}

func UpdateSite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	site := Site{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		URL:         r.FormValue("url"),
	}

	db := helpers.Db

	err := db.QueryRow("UPDATE sites set title=$1, description=$2, url=$3 WHERE id=$4 returning id, created, rating, views", site.Title, site.Description, site.URL, id).Scan(&site.ID, &site.Created, &site.Rating, &site.Views)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Site not found", 404)
		return
	}

	js, _ := json.MarshalIndent(site, "", " ")

	w.Write(js)
}

func DeleteSite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	var deletedID int

	db := helpers.Db

	err := db.QueryRow("DELETE FROM sites WHERE id=$1 returning id", id).Scan(&deletedID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Site not found", 404)
		return
	}

	fmt.Fprint(w, deletedID)
}

func LikeSite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var rating int

	db := helpers.Db

	err := db.QueryRow("UPDATE sites set rating=rating+1 where id=$1 returning rating", id).Scan(&rating)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Site not found", 404)
		return
	}

	fmt.Fprint(w, rating)
}
