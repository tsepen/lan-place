package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetSites(c *gin.Context) {
	fmt.Println("get")
}

func GetOneSite(c *gin.Context) {
	fmt.Println("get")
}

func CreateSite(c *gin.Context) {
	fmt.Println("post")
}

func UpdateSite(c *gin.Context) {
	fmt.Println("update")
}

func DeleteSite(c *gin.Context) {
	fmt.Println("delete")
}

func LikeSite(c *gin.Context) {
	fmt.Println("like")
}
