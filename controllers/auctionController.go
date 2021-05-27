package controllers

import (
	"drouotBack/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	auctionService      = services.NewAuctionService()
	auctionInputService = services.NewCheckAuctionInputService()
)

// GET /auctions
// Get all auctions
func GetAllAuctions(c *gin.Context) {
	auctions := auctionService.FindAuctions()
	c.JSON(http.StatusOK, gin.H{"data": auctions})
}

// GET /auctions/:id
// Get an auction by its id
func GetAuctionById(c *gin.Context) {
	auction, err := auctionService.FindAuction(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": auction})
}

// POST /auctions
// Create a new auction
func AddAuction(c *gin.Context) {
	auction, err1 := auctionInputService.CheckCreateInput(c)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
	}
	newAuction, err2 := auctionService.CreateAuction(auction)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": newAuction})
}

// DELETE /auctions/:id
// Delete an auction
func DeleteAuction(c *gin.Context) {
	deleted, err := auctionService.DeleteAuction(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": deleted})
}

// PUT /auctions/:id
// Update an auction
func UpdateAuction(c *gin.Context) {
	auction, err1 := auctionInputService.CheckUpdateInput(c)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
	}
	newAuction, err2 := auctionService.UpdateAuction(c.Param("id"), auction)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": newAuction})
}

// GET /auctions/articles/:id
// Get all the articles from auction with given id
func GetArticlesAuction(c *gin.Context) {
	articles, err := auctionService.GetArticlesAuction(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": articles})
}
