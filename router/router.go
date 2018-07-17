package router

import (
	"landing-place/handlers"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/sign-in", handlers.SignIn).Methods("POST")
	r.HandleFunc("/api/v1/sign-up", handlers.SignUp).Methods("POST")
	// v1 := router.Group("/api/v1")
	// {
	// 	v1.POST("/sign-in", handlers.SignIn)
	// 	v1.POST("/sign-up", handlers.SignUp)

	// 	v1.GET("/site", handlers.GetSites)
	// 	v1.GET("/site/:id", handlers.GetOneSite)
	// 	v1.POST("/site", handlers.CreateSite)
	// 	v1.POST("/site/:id/like", handlers.LikeSite)
	// 	v1.PUT("/site/:id", handlers.UpdateSite)
	// 	v1.DELETE("/site/:id", handlers.DeleteSite)

	// 	v1.GET("/category", handlers.GetCategories)
	// 	v1.POST("/category", handlers.CreateCategory)
	// 	v1.PUT("/category/:id", handlers.UpdateCategory)
	// 	v1.DELETE("/category/:id", handlers.DeleteCategory)

	// 	v1.GET("/tag", handlers.GetTags)
	// 	v1.POST("/tag", handlers.CreateTag)
	// 	v1.PUT("/tag/:id", handlers.UpdateTag)
	// 	v1.DELETE("/tag/:id", handlers.DeleteTag)
	// }

	return r
}
