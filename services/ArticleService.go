package services

import (
	"restapi/models"
	"restapi/repositories"
)

//IArticleService :
type IArticleService interface {
	GetArticleByID(id uint) (*models.Article, error)
}

//ArticleService :
type ArticleService struct {
	ArticleRepository repositories.IArticleRepository
}

//NewArticleService :
func NewArticleService(articleRepository repositories.IArticleRepository) IArticleService {
	return &ArticleService{
		ArticleRepository: articleRepository,
	}
}

//GetArticleByID :
func (service *ArticleService) GetArticleByID(id uint) (*models.Article, error) {
	return service.ArticleRepository.Get(id)
}
