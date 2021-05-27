package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Auction struct {
	gorm.Model
	Title     string    `gorm:"column:title" json:"title"`
	Category  string    `gorm:"column:category" json:"category"`
	UserID    uint      `gorm:"column:userID" json:"userID"`
	StartDate time.Time `gorm:"column:startDate" json:"startDate"`
	EndDate   time.Time `gorm:"column:endDate" json:"endDate"`
	PhotoURL  string    `gorm:"column:photoURL" json:"photoURL"`
	// Articles  []Article `gorm:"foreignKey:auctionID"` // Does not auto link articles based on the auction id given to article
}

// Update the auction dates using a custom input date format
// The specify parameter should be "start" or "end" depending on which date is to be updated
// func (a *Auction) UpdateDate(tx *gorm.DB, specify string, date string) (err error) {
// 	// format of input date
// 	format := "2006-01-02 15:04:05"
// 	if specify == "start" {
// 		var startDate, _ = time.Parse(format, date)
// 		tx.Model(&a).Update("StartDate", startDate)
// 	} else if specify == "end" {
// 		var endDate, _ = time.Parse(format, date)
// 		tx.Model(&a).Update("EndDate", endDate)
// 	}
// 	return
// }

// Cascade to articles on soft-delete
func (u *Auction) AfterDelete(tx *gorm.DB) (err error) {
	var articles []Article

	DB.Where("auctionId = ?", u.ID).Find(&articles)

	for _, article := range articles {
		DB.Delete(&article)
	}
	return
}

// This allows us to test the format of the supplied information and create an entry in the database.
type CreateAuctionInput struct {
	Title     string    `json:"title" binding:"required"`
	Category  string    `json:"category" binding:"required"`
	UserID    uint      `json:"userId" binding:"required"`
	StartDate time.Time `json:"startDate" binding:"required"`
	EndDate   time.Time `json:"endDate" binding:"required"`
	PhotoURL  string    `json:"photoURL" binding:"required"`
}

// This allows us to test the supplied information and update the entity's information in the database.
type UpdateAuctionInput struct {
	Title     string    `json:"title"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	PhotoURL  string    `json:"photoURL"`
}
