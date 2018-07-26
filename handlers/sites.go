package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"landing-place/models"
	"landing-place/services"
)

func GetSites(w http.ResponseWriter, r *http.Request) {

	rows, err := services.GetSites()
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	sites := make([]*models.Site, 0)

	for rows.Next() {
		site := new(models.Site)

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

	site := new(models.Site)

	row := services.GetOneSite(id)

	err := row.Scan(&site.ID, &site.Title, &site.Description, &site.URL, &site.Created, &site.Rating, &site.Views)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	/* rows, err := db.Query("SELECT s.title,	array_agg(t.title) FILTER (WHERE t.title IS NOT NULL) AS tags FROM sites s LEFT JOIN sitesTags st ON s.id = st.siteid LEFT JOIN tags t ON st.tagid = t.id WHERE s.id = $1 GROUP BY s.title;", id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error", 404)
		return
	}
	fmt.Println(rows) */
	/* 	for rows.Next() {
		js, _ := json.MarshalIndent(rows, "", " ")
		fmt.Println(js)
	} */

	js, _ := json.MarshalIndent(site, "", " ")

	w.Write(js)
}

func CreateSite(w http.ResponseWriter, r *http.Request) {
	site := models.Site{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		URL:         r.FormValue("url"),
	}
	/*
		tags := r.FormValue("tags")
		categories := r.FormValue("categories") */

	err := services.CreateSite(&site)
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

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}

	site := models.Site{
		ID:          id,
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		URL:         r.FormValue("url"),
	}

	rowsAffected, err := services.UpdateSite(&site)
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

	rowsAffected, err := services.DeleteSite(id)
	if err != nil || rowsAffected == 0 {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Site with id %s deleted successfully\n", id)
}

func LikeSite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var rating int

	row := services.LikeSite(id)

	err := row.Scan(&rating)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Site not found", 404)
		return
	}

	fmt.Fprintf(w, "Site with id %s has rating %v\n", id, rating)
}
