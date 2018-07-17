package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	fmt.Println("get")
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
