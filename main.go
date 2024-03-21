package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./web/views/*")
	router.Static("/static","./web/styles")
	router.GET("/content", func(c *gin.Context) {
		c.HTML(http.StatusOK, "content", gin.H{
			"Title": "Posts",
			"Name":  "Niks",
		})
	})

	router.Run(":8080")
}
