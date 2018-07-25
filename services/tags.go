package services

import (
	"errors"
	"landing-place/helpers"
	"landing-place/models"
)

func CreateTag(tag *models.Tag) error {
	db := helpers.Db

	err := db.QueryRow("INSERT INTO tags(title) VALUES($1) returning id", tag.Title).Scan(&tag.ID)
	if err != nil {
		return errors.New("Tag title in database")
	}
	return nil
}
