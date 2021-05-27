package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title           string `gorm:"column:title" json:"title"`
	Description     string `gorm:"column:description" json:"description"`
	AuctionId       uint   `gorm:"column:auctionID" json:"auctionID"`
	Estimation      uint   `gorm:"column:estimation" json:"estimation"`
	InitialOffering uint   `gorm:"column:initialOffering" json:"initialOffering"`
	PhotoPath       string `gorm:"column:photoUrl" json:"photoURL"`
}

// Cascade to bids on soft-delete
func (u *Article) AfterDelete(tx *gorm.DB) (err error) {
	var bids []Bid

	DB.Where("articleId = ?", u.ID).Find(&bids)

	for _, bid := range bids {
		DB.Delete(&bid)
	}
	return
}

// This allows us to test the format of the supplied information and create an entry in the database.
type CreateArticleInput struct {
	Title           string `json:"title" binding:"required"`
	Description     string `json:"description" binding:"required"`
	AuctionId       uint   `json:"auctionID" binding:"required"`
	Estimation      uint   `json:"estimation" binding:"required"`
	InitialOffering uint   `json:"initialOffering" binding:"required"`
	PhotoPath       string `json:"photoURL" binding:"required"`
}

// This allows us to test the supplied information and update the entity's information in the database.
type UpdateArticleInput struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	Estimation      uint   `json:"estimation"`
	InitialOffering uint   `json:"initialOffering"`
	PhotoPath       string `json:"photoURL"`
}
