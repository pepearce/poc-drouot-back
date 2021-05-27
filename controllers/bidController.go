package controllers

import (
	"drouotBack/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	bidService      = services.NewBidService()
	bidInputService = services.NewCheckBidInputService()
)

// GET /bids
// Get all bids
func GetAllbids(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": bidService.FindBids()})
}

// GET /bids/:id
// Get a bid by his id
func GetBidById(c *gin.Context) {
	bid, err := bidService.FindBid(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bid})
}

// POST /bids
// Add a bid
func AddBid(c *gin.Context) {
	bid, err := bidInputService.CheckCreateInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBid, err := bidService.CreateBid(bid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": newBid})
}

// DELETE /bids/:id
// Delete bid
func DeleteBid(c *gin.Context) {
	deleted, err := bidService.DeleteBid(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": deleted})
}
