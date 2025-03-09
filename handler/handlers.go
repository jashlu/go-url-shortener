package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jashlu/go-url-shortener/shortener"
	"github.com/jashlu/go-url-shortener/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

// get the creation request body, parse it, and extract the initial long url and userID
func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	//we get the creation request body, parse it and extract the longId and UserId
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//use generateshortlink to create the shortened url
	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	//we save it to our redis db
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808/"
	//send back success message
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
