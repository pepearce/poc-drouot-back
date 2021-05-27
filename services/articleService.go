package services

import (
	"drouotBack/models"
)

type ArticleService interface {
	FindArticles() []models.Article
	FindArticle(string) (models.Article, error)
	CreateArticle(models.Article) (models.Article, error)
	DeleteArticle(string) (bool, error)
	UpdateArticle(string, models.UpdateArticleInput) (models.Article, error)
	GetBidsArticle(string) ([]models.Bid, error)
	GetHighestBidArticle(string) (models.Bid, error)
}

type articleService struct {
}

func NewArticleService() ArticleService {
	return &articleService{}
}

// returns all articles in DB
func (service *articleService) FindArticles() []models.Article {
	var articles []models.Article
	models.DB.Find(&articles)
	return articles
}

// Finds an article by id
func (service *articleService) FindArticle(id string) (models.Article, error) {
	var article models.Article

	if err := models.DB.Where("id = ?", id).First(&article).Error; err != nil {
		return article, err
	}
	return article, nil
}

// Create a new article
func (service *articleService) CreateArticle(article models.Article) (models.Article, error) {

	if err := models.DB.Create(&article).Error; err != nil {
		return article, err
	}
	return article, nil
}

// Update an article
func (service *articleService) UpdateArticle(id string, articleUpdate models.UpdateArticleInput) (models.Article, error) {
	var article models.Article

	if err := models.DB.Where("id = ?", id).First(&article).Error; err != nil {
		return article, err
	}

	if err := models.DB.Model(&article).Updates(articleUpdate).Error; err != nil {
		return article, err
	}
	return article, nil
}

// Delete an article by id
func (service *articleService) DeleteArticle(id string) (bool, error) {
	var article models.Article

	if err := models.DB.Where("id = ?", id).First(&article).Error; err != nil {
		return false, err
	}

	if err := models.DB.Delete(&article).Error; err != nil {
		return false, err
	}

	return true, nil
}

// Get all bids for article with given id
func (service *articleService) GetBidsArticle(id string) ([]models.Bid, error) {
	var bids []models.Bid

	if err := models.DB.Where("articleId = ?", id).Find(&bids).Error; err != nil {
		return bids, err
	}

	return bids, nil
}

// Get highest bid for article with given id
func (service *articleService) GetHighestBidArticle(id string) (models.Bid, error) {
	var bid models.Bid

	if err := models.DB.Where("articleId = ? and bidAmount = (select max(bidAmount) from bids where articleId = ? )", id, id).First(&bid).Error; err != nil {
		if err.Error() == "record not found" {
			//if no bid exists, returns an empty bid without an error
			return bid, nil
		} else {
			// if another error occurs, returns empty bid and error
			return bid, err
		}
	}
	return bid, nil
}
