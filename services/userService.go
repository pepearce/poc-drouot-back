package services

import (
	"drouotBack/models"
	"drouotBack/security"
	"errors"

	"github.com/google/uuid"
)

// All errors occuring during service operations are passed down to the controller.
type UserService interface {
	FindUsers() []models.User
	FindUser(string) (models.User, error)
	CreateUser(models.User) (models.User, error)
	DeleteUser(string) (bool, error)
	UpdateUser(string, models.UpdateUserInput) (models.User, error)
	SignInUser(models.SignInUserInput) (models.User, error)
}

type userService struct{}

func NewUserService() UserService {

	return &userService{}
}

// returns all users in DB
func (service *userService) FindUsers() []models.User {
	var users []models.User
	models.DB.Find(&users)
	return users
}

// Create a new user

func (service *userService) CreateUser(user models.User) (models.User, error) {
	user.UUID = uuid.NewString()
	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPassword)
	if err := models.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Finds a user
func (service *userService) FindUser(userId string) (models.User, error) {
	var user models.User

	if err := models.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Update a user
func (service *userService) UpdateUser(userId string, userInput models.UpdateUserInput) (models.User, error) {
	var user models.User

	if err := models.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return user, err
	}

	if userInput.FirstName != "" && userInput.FirstName != user.FirstName {
		user.FirstName = userInput.FirstName
	}
	if userInput.LastName != "" && userInput.LastName != user.LastName {
		user.LastName = userInput.LastName
	}
	if userInput.Address != "" && userInput.Address != user.Address {
		user.Address = userInput.Address
	}

	if err := models.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Delete a user using the gin context
func (service *userService) DeleteUser(userId string) (bool, error) {
	var user models.User

	if err := models.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return false, err
	}

	models.DB.Delete(&user)
	return true, nil
}

// SignInUser /signIn
func (service *userService) SignInUser(signInInput models.SignInUserInput) (models.User, error) {
	var user models.User

	if err := models.DB.Where("email = ?", signInInput.Email).First(&user).Error; err != nil {
		return user, err
	}
	if err := security.VerifyPassword(user.Password, signInInput.Password); err != nil {
		return user, errors.New("incorrect password")
	} else {
		return user, nil
	}
}
