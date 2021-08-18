package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zkhorozianbc/url-shortener/shortener"
	"github.com/zkhorozianbc/url-shortener/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest

	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortUrl(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808/"

	c.JSON(200, gin.H{
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {

	/// Implementation to be added

	shortUrl := c.Param("shortUrl")
	originalUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, originalUrl)

}
