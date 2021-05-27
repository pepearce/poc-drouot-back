package services

import (
	"drouotBack/models"
)

type BidService interface {
	FindBids() []models.Bid
	FindBid(string) (models.Bid, error)
	CreateBid(models.Bid) (models.Bid, error)
	DeleteBid(string) (bool, error)
}

type bidService struct {
}

func NewBidService() BidService {
	return &bidService{}
}

// GET /bids
// returns all bids in DB
func (bidService *bidService) FindBids() []models.Bid {
	var bids []models.Bid
	models.DB.Find(&bids)
	return bids
}

// GET /bids/:id
// Finds an bid by id
func (bidService *bidService) FindBid(id string) (models.Bid, error) {
	var bid models.Bid

	if err := models.DB.Where("id = ?", id).First(&bid).Error; err != nil {
		return bid, err
	}
	return bid, nil
}

// POST /bids
// Create a new bid
func (bidService *bidService) CreateBid(bid models.Bid) (models.Bid, error) {

	if err := models.DB.Create(&bid).Error; err != nil {
		return bid, err
	}

	return bid, nil
}

// DELETE /bids/id
// Delete a bid by id
func (bidService *bidService) DeleteBid(id string) (bool, error) {
	var bid models.Bid

	if err := models.DB.Where("id = ?", id).First(&bid).Error; err != nil {
		return false, err
	}

	if err := models.DB.Delete(&bid).Error; err != nil {
		return false, err
	}
	return true, nil
}
