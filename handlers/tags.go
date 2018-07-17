package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	fmt.Println("get")
}

func CreateTag(c *gin.Context) {
	fmt.Println("post")
}

func UpdateTag(c *gin.Context) {
	fmt.Println("update")
}

func DeleteTag(c *gin.Context) {
	fmt.Println("delete")
}
