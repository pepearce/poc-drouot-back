package controllers

import (
	"drouotBack/services"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	service      = services.NewArticleService()
	inputService = services.NewCheckArticleInputService()
)

// GET /articles
// Get all articles
func GetAllArticles(c *gin.Context) {
	session := sessions.Default(c)
	log.Println(session.Get("session-id"))
	articles := service.FindArticles()
	c.JSON(http.StatusOK, gin.H{"data": articles})
}

// GET /articles/:id
// Get an article by its id
func GetArticleById(c *gin.Context) {
	article, err := service.FindArticle(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": article})
}

// POST /articles
// Create a new article
func AddArticle(c *gin.Context) {
	article, err1 := inputService.CheckCreateInput(c)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
	}
	newArticle, err2 := service.CreateArticle(article)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": newArticle})
}

// DELETE /articles/:id
// Delete an article
func DeleteArticle(c *gin.Context) {
	deleted, err := service.DeleteArticle(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": deleted})
}

// PUT /articles/:id
// Update an article
func UpdateArticle(c *gin.Context) {
	article, err1 := inputService.CheckUpdateInput(c)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
	}
	newArticle, err2 := service.UpdateArticle(c.Param("id"), article)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": newArticle})
}

// GET /articles/bids/:id
// Get all bids for article with given id
func GetBidsArticle(c *gin.Context) {
	bids, err := service.GetBidsArticle(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": bids})
}

// GET /articles/bid/:id
// Get highest bid for article with given id
func GetHighestBidArticle(c *gin.Context) {
	bid, err := service.GetHighestBidArticle(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": bid})
}
