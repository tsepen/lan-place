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

	categories, err := db.Query("CREATE TABLE IF NOT EXISTS categories (id SERIAL PRIMARY KEY,title CHARACTER VARYING(40),UNIQUE(title));")
	if err != nil {
		log.Println(err)
	}

	tags, err := db.Query("CREATE TABLE IF NOT EXISTS tags (id SERIAL PRIMARY KEY,title CHARACTER VARYING(40),UNIQUE(title));")
	if err != nil {
		log.Println(err)
	}

	users, err := db.Query("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY,name CHARACTER VARYING(40),email CHARACTER VARYING(40),password text,UNIQUE(email));")
	if err != nil {
		log.Println(err)
	}

	sites, err := db.Query("CREATE TABLE IF NOT EXISTS sites (id SERIAL PRIMARY KEY,title CHARACTER VARYING(40),description text,url CHARACTER VARYING(40),    created DATE default NOW(),rating INTEGER default 0,views INTEGER default 0, categories CHARACTER VARYING(40)[],tags CHARACTER VARYING(40)[],    UNIQUE(url));")
	if err != nil {
		log.Println(err)
	}

	defer categories.Close()
	defer users.Close()
	defer sites.Close()
	defer tags.Close()
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
