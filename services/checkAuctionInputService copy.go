package services

import (
	"drouotBack/models"

	"github.com/gin-gonic/gin"
)

type CheckAuctionInputService interface {
	CheckCreateInput(*gin.Context) (models.Auction, error)
	CheckUpdateInput(*gin.Context) (models.UpdateAuctionInput, error)
}

type checkAuctionInputService struct{}

func NewCheckAuctionInputService() CheckAuctionInputService {
	return &checkAuctionInputService{}
}

func (service *checkAuctionInputService) CheckCreateInput(c *gin.Context) (models.Auction, error) {
	var input models.CreateAuctionInput
	var newAuction models.Auction

	if err := c.ShouldBindJSON(&input); err != nil {
		return newAuction, err
	}
	newAuction = models.Auction{Title: input.Title,
		Category:  input.Category,
		UserID:    input.UserID,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		PhotoURL:  input.PhotoURL}

	return newAuction, nil
}
func (service *checkAuctionInputService) CheckUpdateInput(c *gin.Context) (models.UpdateAuctionInput, error) {
	var input models.UpdateAuctionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return input, err
	}
	return input, nil
}
