package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Bid struct {
	gorm.Model
	BidDate   time.Time `gorm:"column:bidDate" json:"bidDate"`
	UserID    uint      `gorm:"column:userID" json:"userId"`
	ArticleID uint      `gorm:"column:articleID" json:"articleId"`
	BidAmount float32   `gorm:"column:bidAmount" json:"bidAmount"`
}

// This allows us to test the format of the supplied information and create an entry in the database.
type CreateBidInput struct {
	BidDate   time.Time `json:"bidDate" binding:"required"`
	UserID    uint      `json:"userId" binding:"required"`
	ArticleID uint      `json:"articleId" binding:"required"`
	BidAmount float32   `json:"bidAmount" binding:"required"`
}

// // This allows us to test the supplied information and update the entity's information in the database.
// type UpdateBidInput struct {
// 	BidDate string `json:"bidDate" binding:"requirer"`
// 	UserID     uint `json:"userId" binding:"required"`
// 	ArticleID uint `json:"articleId" binding:"required"`
// 	BidAmount   float32 `json:"bidAmount" binding:"required"`
// }

// func (bid *Bid) AfterCreate(tx *gorm.DB) (err error) {
// 	var article Article
// 	DB.Where("id = ?", bid.ArticleID).First(&article)
// 	article.HighestBidId = bid.ID
// 	DB.Save(&article)
// 	return
// }
