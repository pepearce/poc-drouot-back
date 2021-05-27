package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UUID      string `gorm:"column:uuid" json:"sessionId"`
	FirstName string `gorm:"column:firstName" json:"firstName"`
	LastName  string `gorm:"column:lastName" json:"lastName"`
	Address   string `gorm:"column:address" json:"address"`
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"password"`
	Role      string `gorm:"column:role;default:'user'" json:"role"`
}

// Cascade to bids on soft-delete
func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	var bids []Bid

	DB.Where("userId = ?", u.ID).Find(&bids)

	for _, bid := range bids {
		DB.Delete(&bid)
	}
	return
}

// This allows us to test the supplied information.
type CreateUserInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Address   string `json:"address" `
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role" binding:"required"`
}

// This allows us to test the supplied information.
type UpdateUserInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
	// Email     string `json:"email"`
	// Password  string `json:"password"`
}

// This allows us to test the supplied information.
type SignInUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
