package router

import (
	"landing-place/handlers"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/sign-in", handlers.SignIn).Methods("POST")
	v1.HandleFunc("/sign-up", handlers.SignUp).Methods("POST")
	v1.HandleFunc("/sign-out", handlers.SignOut).Methods("GET")

	v1.HandleFunc("/site", handlers.GetSites).Methods("GET")
	v1.HandleFunc("/site", handlers.Auth(handlers.CreateSite)).Methods("POST")
	v1.HandleFunc("/site/{id}", handlers.GetOneSite).Methods("GET")
	v1.HandleFunc("/site/{id}", handlers.Auth(handlers.UpdateSite)).Methods("PUT")
	v1.HandleFunc("/site/{id}", handlers.Auth(handlers.DeleteSite)).Methods("DELETE")
	v1.HandleFunc("/site/{id}/like", handlers.LikeSite).Methods("PUT")

	v1.HandleFunc("/tag", handlers.GetTags).Methods("GET")
	v1.HandleFunc("/tag", handlers.Auth(handlers.CreateTag)).Methods("POST")
	v1.HandleFunc("/tag/{id}", handlers.GetOneTag).Methods("GET")
	v1.HandleFunc("/tag/{id}", handlers.Auth(handlers.UpdateTag)).Methods("PUT")
	v1.HandleFunc("/tag/{id}", handlers.Auth(handlers.DeleteTag)).Methods("DELETE")

	v1.HandleFunc("/category", handlers.GetCategories).Methods("GET")
	v1.HandleFunc("/category", handlers.Auth(handlers.CreateCategory)).Methods("POST")
	v1.HandleFunc("/category/{id}", handlers.GetOneCategory).Methods("GET")
	v1.HandleFunc("/category/{id}", handlers.Auth(handlers.UpdateCategory)).Methods("PUT")
	v1.HandleFunc("/category/{id}", handlers.Auth(handlers.DeleteCategory)).Methods("DELETE")

	return r
}
