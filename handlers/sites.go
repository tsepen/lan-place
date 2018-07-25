package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"landing-place/helpers"
)

type tag struct {
	ID    int    `json:"id"`
	title string `json:"title"`
}

type category struct {
	Siteid     int    `json:"siteId"`
	CategoryID string `json:"categoryId"`
}

type Site struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	URL         string      `json:"url"`
	Created     string      `json:"createdAt"`
	Rating      int         `json:"rating"`
	Views       int         `json:"views"`
	Tags        []*tag      `json:"tags"`
	Categories  []*category `json:"categories"`
}

func GetSites(w http.ResponseWriter, r *http.Request) {
	db := helpers.Db

	rows, err := db.Query("SELECT * FROM sites")
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

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

	row := db.QueryRow("UPDATE sites SET views=views+1 WHERE id=$1 returning *", id)
	err := row.Scan(&site.ID, &site.Title, &site.Description, &site.URL, &site.Created, &site.Rating, &site.Views)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	rows, err := db.Query("SELECT s.title,	array_agg(t.title) FILTER (WHERE t.title IS NOT NULL) AS tags FROM sites s LEFT JOIN sitesTags st ON s.id = st.siteid LEFT JOIN tags t ON st.tagid = t.id WHERE s.id = $1 GROUP BY s.title;", id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error", 404)
		return
	}
	fmt.Println(rows)
	/* 	for rows.Next() {
		js, _ := json.MarshalIndent(rows, "", " ")
		fmt.Println(js)
	} */

	js, _ := json.MarshalIndent(site, "", " ")

	w.Write(js)
}

func CreateSite(w http.ResponseWriter, r *http.Request) {
	site := Site{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		URL:         r.FormValue("url"),
	}
	/*
		tags := r.FormValue("tags")
		categories := r.FormValue("categories") */

	db := helpers.Db

	err := db.QueryRow("INSERT INTO sites(title, description, url) VALUES($1, $2, $3) returning id, created, rating, views", site.Title, site.Description, site.URL).Scan(&site.ID, &site.Created, &site.Rating, &site.Views)
	if err != nil {
		http.Error(w, "This site in database", 400)
		return
	}

	/* var validID = regexp.MustCompile(`^[0-9]$`)

	for _, value := range tags {
		valid := validID.MatchString(string(value))

		if valid == true {

			tagID := string(value)

			rows, err := db.Query("INSERT INTO sitesTags(siteId, tagId) VALUES($1, $2) returning *", site.ID, tagID)
			if err != nil {
				fmt.Println(err)
				return
			}

			for rows.Next() {
								tag := new(tag)

				   				rows.Scan(&tag.Siteid, &tag.TagID)
				   				site.Tags = append(site.Tags, tag)
			}
		}
	}

	for _, value := range categories {
		valid := validID.MatchString(string(value))

		if valid == true {

			categoryID := string(value)

			rows, err := db.Query("INSERT INTO sitesCategories(siteId, categoryId) VALUES($1, $2) returning *", site.ID, categoryID)
			if err != nil {
				fmt.Println(err)
				return
			}

			for rows.Next() {
								category := new(category)

				   				rows.Scan(&category.Siteid, &category.CategoryID)
				   				site.Categories = append(site.Categories, category)
			}
		}
	}
	*/
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

	result, err := db.Exec("UPDATE sites set title=$1, description=$2, url=$3 WHERE id=$4 returning id, created, rating, views", site.Title, site.Description, site.URL, id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, http.StatusText(500), 500)
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
