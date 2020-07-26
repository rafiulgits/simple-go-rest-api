package repositories

import (
	"fmt"
	"restapi/conn"
	"restapi/models"
)

//IArticleRepository :
type IArticleRepository interface {
	Get(id uint) (*models.Article, error)
	GetAll() ([]*models.Article, error)
}

//ArticleRepository :
type ArticleRepository struct {
	*BaseRepository
}

//NewArticleRepository :
func NewArticleRepository(db *conn.DB) IArticleRepository {
	fmt.Println("injected")
	return &ArticleRepository{
		&BaseRepository{
			db: db.Table("Articles"),
		},
	}
}

//Get : Get an article by Id
func (repo *ArticleRepository) Get(id uint) (*models.Article, error) {
	return nil, nil
}

//GetAll : Getll all articles
func (repo *ArticleRepository) GetAll() ([]*models.Article, error) {
	return []*models.Article{}, nil
}
