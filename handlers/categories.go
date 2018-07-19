package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("get")

	// session, _ := store.Get(r, "cookie-name")

	// // Check if user is authenticated
	// if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
	// 	http.Error(w, "Forbidden", http.StatusForbidden)
	// 	return
	// }

	// // Print secret message
	// fmt.Fprintln(w, "The cake is a lie!")

}

func CreateCategory(c *gin.Context) {
	fmt.Println("post")
}

func UpdateCategory(c *gin.Context) {
	fmt.Println("update")
}

func DeleteCategory(c *gin.Context) {
	fmt.Println("delete")
}
