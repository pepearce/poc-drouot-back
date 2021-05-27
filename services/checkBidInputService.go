package services

import (
	"drouotBack/models"

	"github.com/gin-gonic/gin"
)

type CheckBidInputService interface {
	CheckCreateInput(*gin.Context) (models.Bid, error)
}

type checkBidInputService struct{}

func NewCheckBidInputService() CheckBidInputService {
	return &checkBidInputService{}
}

func (service *checkBidInputService) CheckCreateInput(c *gin.Context) (models.Bid, error) {
	var input models.CreateBidInput
	var newBid models.Bid

	if err := c.ShouldBindJSON(&input); err != nil {
		return newBid, err
	}
	newBid = models.Bid{BidDate: input.BidDate,
		UserID:    input.UserID,
		ArticleID: input.ArticleID,
		BidAmount: input.BidAmount}

	return newBid, nil
}
