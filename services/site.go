package services

import (
	"database/sql"
	"fmt"
	"landing-place/helpers"
	"landing-place/models"
)

func GetSites() (*sql.Rows, error) {
	db := helpers.Db

	rows, err := db.Query("SELECT * FROM sites")

	return rows, err
}

func GetOneSite(id string) *sql.Row {
	db := helpers.Db

	row := db.QueryRow("UPDATE sites SET views=views+1 WHERE id=$1 returning *", id)

	return row
}

func CreateSite(site *models.Site) error {
	db := helpers.Db

	err := db.QueryRow("INSERT INTO sites(title, description, url) VALUES($1, $2, $3) returning id, created, rating, views", &site.Title, &site.Description, &site.URL).Scan(&site.ID, &site.Created, &site.Rating, &site.Views)

	return err
}

func UpdateSite(site *models.Site) (int64, error) {
	db := helpers.Db

	rows, err := db.Exec("UPDATE sites set title=$1, description=$2, url=$3 WHERE id=$4 returning id, created, rating, views", &site.Title, &site.Description, &site.URL, &site.ID)
	if err != nil {
		fmt.Println(err)
	}

	rowsAffected, err := rows.RowsAffected()

	return rowsAffected, err
}

func DeleteSite(id string) (int64, error) {
	db := helpers.Db

	rows, err := db.Exec("DELETE FROM sites WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
	}

	rowsAffected, err := rows.RowsAffected()

	return rowsAffected, err
}

func LikeSite(id string) *sql.Row {
	db := helpers.Db

	row := db.QueryRow("UPDATE sites set rating=rating+1 where id=$1 returning rating", id)

	return row
}
