package main

import (
	"fmt"
	"landing-place/helpers"
	"landing-place/router"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)

func init() {
	helpers.ConnectDB()
}

func main() {
	r := router.New()
	// r.Use(cors.Default())

	fmt.Println("Server is listening 8080...")
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r))
}
