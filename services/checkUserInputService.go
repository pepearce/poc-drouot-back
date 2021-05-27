package services

import (
	"drouotBack/models"

	"github.com/gin-gonic/gin"
)

type CheckUserInputService interface {
	CheckCreateInput(*gin.Context) (models.User, error)
	CheckUpdateInput(*gin.Context) (models.UpdateUserInput, error)
	CheckSignInInput(*gin.Context) (models.SignInUserInput, error)
}

type checkUserInputService struct{}

func NewCheckUserInputService() CheckUserInputService {
	return &checkUserInputService{}
}

func (service *checkUserInputService) CheckCreateInput(c *gin.Context) (models.User, error) {
	var input models.CreateUserInput
	var newUser models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		return newUser, err
	}
	newUser = models.User{FirstName: input.FirstName,
		LastName: input.LastName,
		Email:    input.Email,
		Address:  input.Address,
		Password: input.Password,
		Role:     input.Role}

	return newUser, nil
}
func (service *checkUserInputService) CheckUpdateInput(c *gin.Context) (models.UpdateUserInput, error) {
	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return input, err
	}
	return input, nil
}
func (service *checkUserInputService) CheckSignInInput(c *gin.Context) (models.SignInUserInput, error) {
	var input models.SignInUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return input, err
	}
	return input, nil
}
