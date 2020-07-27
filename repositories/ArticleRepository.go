package repositories

import (
	"restapi/conn"
	"restapi/models"
)

//IArticleRepository :
type IArticleRepository interface {
	Get(id uint) (*models.Article, error)
	GetAll() ([]*models.Article, error)
	Create(a *models.Article) (*models.Article, error)
}

//ArticleRepository :
type ArticleRepository struct {
	*BaseRepository
}

//NewArticleRepository :
func NewArticleRepository(db *conn.DB) IArticleRepository {
	return &ArticleRepository{
		&BaseRepository{
			db: db.Table("Articles"),
		},
	}
}

//Get : Get an article by Id
func (repo *ArticleRepository) Get(id uint) (*models.Article, error) {
	var article models.Article
	err := repo.db.First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil

}

//GetAll : Getll all articles
func (repo *ArticleRepository) GetAll() ([]*models.Article, error) {
	var articles []*models.Article
	err := repo.db.Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

//Create :
func (repo *ArticleRepository) Create(a *models.Article) (*models.Article, error) {
	var err error
	if repo.db.NewRecord(a) {
		repo.db.Create(&a)
		if !repo.db.NewRecord(a) {
			return a, nil
		}
		return nil, err
	}
	return nil, err
}
