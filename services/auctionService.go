package services

import (
	"drouotBack/models"
)

type AuctionService interface {
	FindAuctions() []models.Auction
	FindAuction(string) (models.Auction, error)
	CreateAuction(models.Auction) (models.Auction, error)
	DeleteAuction(string) (bool, error)
	UpdateAuction(string, models.UpdateAuctionInput) (models.Auction, error)
	GetArticlesAuction(string) ([]models.Article, error)
}

type auctionService struct {
}

func NewAuctionService() AuctionService {
	return &auctionService{}
}

// returns all auctions in DB
func (service *auctionService) FindAuctions() []models.Auction {
	var auctions []models.Auction
	models.DB.Find(&auctions)
	return auctions
}

// Finds an auction by id
func (service *auctionService) FindAuction(id string) (models.Auction, error) {
	var auction models.Auction

	if err := models.DB.Where("id = ?", id).First(&auction).Error; err != nil {
		return auction, err
	}
	return auction, nil
}

// Create a new auction
func (service *auctionService) CreateAuction(auction models.Auction) (models.Auction, error) {

	if err := models.DB.Create(&auction).Error; err != nil {
		return auction, err
	}
	return auction, nil
}

// Update an auction
func (service *auctionService) UpdateAuction(id string, auctionUpdate models.UpdateAuctionInput) (models.Auction, error) {
	var auction models.Auction

	if err := models.DB.Where("id = ?", id).First(&auction).Error; err != nil {
		return auction, err
	}

	if err := models.DB.Model(&auction).Updates(auctionUpdate).Error; err != nil {
		return auction, err
	}
	return auction, nil
}

// Delete an auction by id
func (service *auctionService) DeleteAuction(id string) (bool, error) {
	var auction models.Auction

	if err := models.DB.Where("id = ?", id).First(&auction).Error; err != nil {
		return false, err
	}

	if err := models.DB.Delete(&auction).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (service *auctionService) GetArticlesAuction(id string) ([]models.Article, error) {
	var articles []models.Article

	if err := models.DB.Where("auctionId = ?", id).Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
}
