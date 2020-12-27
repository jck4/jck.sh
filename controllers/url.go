package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jck4/jck.sh/models"
)

type CreateurlInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateurlInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GET /url
// Find all urls
func FindURLs(c *gin.Context) {
	var url []models.URLCollection
	models.DB.Find(&url)

	c.JSON(http.StatusOK, gin.H{"data": url})
}

// POST /url
// Create new url
func CreateURL(c *gin.Context) {
	// Validate input
	var input models.URLCollectionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create url
	url := models.URLCollection{ActualURL: input.ActualURL, ShortURL: input.ShortURL}
	models.DB.Create(&url)

	c.JSON(http.StatusOK, gin.H{"data": url})
}
