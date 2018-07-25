package helpers

import (
	"database/sql"
	"landing-place/config"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func createTables() {
	db := Db

	_, categoriesErr := db.Exec("CREATE TABLE IF NOT EXISTS categories (id SERIAL PRIMARY KEY,title CHARACTER VARYING(40),UNIQUE(title));")
	if categoriesErr != nil {
		log.Println(categoriesErr)
	}

	_, tagsErr := db.Exec("CREATE TABLE IF NOT EXISTS tags (id SERIAL PRIMARY KEY,title CHARACTER VARYING(40),UNIQUE(title));")
	if tagsErr != nil {
		log.Println(tagsErr)
	}

	_, usersErr := db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY,name CHARACTER VARYING(40),email CHARACTER VARYING(40),password text,UNIQUE(email));")
	if usersErr != nil {
		log.Println(usersErr)
	}

	_, sitesErr := db.Exec("CREATE TABLE IF NOT EXISTS sites (id SERIAL PRIMARY KEY,title CHARACTER VARYING(40),description text,url CHARACTER VARYING(40),    created DATE default NOW(),rating INTEGER default 0,views INTEGER default 0, UNIQUE(url));")
	if sitesErr != nil {
		log.Println(sitesErr)
	}

	_, sitesCategoriesErr := db.Exec("CREATE TABLE IF NOT EXISTS sitesCategories (siteId integer not null references sites(id),categoryId integer not null references categories(id));")
	if sitesCategoriesErr != nil {
		log.Println(sitesCategoriesErr)
	}

	_, sitesTagsErr := db.Exec("CREATE TABLE IF NOT EXISTS sitesTags (siteId integer not null references sites(id),tagId integer not null references tags(id));")
	if sitesTagsErr != nil {
		log.Println(sitesTagsErr)
	}
}

func ConnectDB() {
	db, err := sql.Open("postgres", config.Database)
	if err != nil {
		panic(err)
	}

	Db = db

	createTables()
	println("Connect to Database")

}
