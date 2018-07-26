package services

import (
	"database/sql"
	"errors"
	"fmt"
	"landing-place/helpers"
	"landing-place/models"
)

func GetTag() (*sql.Rows, error) {
	db := helpers.Db

	rows, err := db.Query("SELECT * FROM tags")

	return rows, err
}

func GetOneTag(id int) *sql.Row {
	db := helpers.Db

	row := db.QueryRow("SELECT * FROM tags WHERE id=$1", id)

	return row
}

func CreateTag(tag *models.Tag) error {
	db := helpers.Db

	err := db.QueryRow("INSERT INTO tags(title) VALUES($1) returning id", tag.Title).Scan(&tag.ID)
	if err != nil {
		return errors.New("Tag title in database")
	}
	return nil
}

func UpdateTag(tag *models.Tag) (int64, error) {
	db := helpers.Db

	rows, err := db.Exec("UPDATE tags SET title = $1 WHERE id = $2", &tag.Title, &tag.ID)
	if err != nil {
		fmt.Println(err)
	}

	rowsAffected, err := rows.RowsAffected()

	return rowsAffected, err
}

func DeleteTag(id string) (int64, error) {
	db := helpers.Db

	rows, err := db.Exec("DELETE FROM tags WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
	}

	rowsAffected, err := rows.RowsAffected()

	return rowsAffected, err
}
