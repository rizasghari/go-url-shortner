package handler

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rizasghari/go-url-shortener/shortener"
	"github.com/rizasghari/go-url-shortener/store"
)

type UrlCreationRequest struct {
	OriginalUrl string `json:"originalUrl" binding:"required"`
	UserId  string `json:"userId" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if !isValidURL(creationRequest.OriginalUrl) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The URL is not in a valid format"})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.OriginalUrl, creationRequest.UserId)

	host := "http://localhost:9808/"

	if store.CheckIfExixsts(shortUrl) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Short url already exists",
			"short_url": host + shortUrl,
		})
		return
	}

	store.SaveUrlMapping(shortUrl, creationRequest.OriginalUrl, creationRequest.UserId)
	c.JSON(http.StatusCreated, gin.H{
		"message":   "Short url created successfully ",
		"short_url": host +shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}

func isValidURL(inputURL string) bool {
	parsedURL, err := url.Parse(inputURL)
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != "" && len(strings.Trim(inputURL, " ")) > 0
}
