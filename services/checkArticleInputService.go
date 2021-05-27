package services

import (
	"drouotBack/models"

	"github.com/gin-gonic/gin"
)

type CheckArticleInputService interface {
	CheckCreateInput(*gin.Context) (models.Article, error)
	CheckUpdateInput(*gin.Context) (models.UpdateArticleInput, error)
}

type checkArticleInputService struct{}

func NewCheckArticleInputService() CheckArticleInputService {
	return &checkArticleInputService{}
}

func (service *checkArticleInputService) CheckCreateInput(c *gin.Context) (models.Article, error) {
	var input models.CreateArticleInput
	var newArticle models.Article

	if err := c.ShouldBindJSON(&input); err != nil {
		return newArticle, err
	}
	newArticle = models.Article{Title: input.Title,
		Description:     input.Description,
		AuctionId:       input.AuctionId,
		Estimation:      input.Estimation,
		InitialOffering: input.InitialOffering,
		PhotoPath:       input.PhotoPath}

	return newArticle, nil
}
func (service *checkArticleInputService) CheckUpdateInput(c *gin.Context) (models.UpdateArticleInput, error) {
	var input models.UpdateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return input, err
	}
	return input, nil
}
