package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/zkhorozianbc/url-shortener/handlers"
	"github.com/zkhorozianbc/url-shortener/store"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Mr. Shrimps!",
		})
	})

	r.POST("/create", func(c *gin.Context) {
		handlers.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handlers.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	fmt.Println("Initializing server")
	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server. Error: %v", err))
	}

}
