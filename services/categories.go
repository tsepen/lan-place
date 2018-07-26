package services

import (
	"database/sql"
	"errors"
	"fmt"
	"landing-place/helpers"
	"landing-place/models"
)

func GetCategory() (*sql.Rows, error) {
	db := helpers.Db

	rows, err := db.Query("SELECT * FROM categories")

	return rows, err
}

func GetOneCategory(id string) *sql.Row {
	db := helpers.Db

	row := db.QueryRow("SELECT * FROM categories WHERE id=$1", id)

	return row
}

func CreateCategory(category *models.Category) error {
	db := helpers.Db

	err := db.QueryRow("INSERT INTO categories(title) VALUES($1) returning id", category.Title).Scan(&category.ID)
	if err != nil {
		return errors.New("Category title in database")
	}
	return nil
}

func UpdateCategory(category *models.Category) (int64, error) {
	db := helpers.Db

	rows, err := db.Exec("UPDATE categories SET title = $1 WHERE id = $2", &category.Title, &category.ID)
	if err != nil {
		fmt.Println(err)
	}

	rowsAffected, err := rows.RowsAffected()

	return rowsAffected, err
}

func DeleteCategory(id string) (int64, error) {
	db := helpers.Db

	rows, err := db.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
	}

	rowsAffected, err := rows.RowsAffected()

	return rowsAffected, err
}
