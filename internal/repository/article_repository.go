package repository

import (
	"portfolio-backend/internal/domain/entity"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func (r *ArticleRepository) Create(article *entity.Article) error {
	return r.DB.Create(article).Error
}

func (r *ArticleRepository) FindAll() ([]entity.Article, error) {
	var articles []entity.Article
	err := r.DB.Order("created_at DESC").Find(&articles).Error
	return articles, err
}

func (r *ArticleRepository) FindByID(id string) (*entity.Article, error) {
	var article entity.Article
	err := r.DB.First(&article, "id = ?", id).Error
	return &article, err
}

func (r *ArticleRepository) Update(article *entity.Article) error {
	return r.DB.Save(article).Error
}

func (r *ArticleRepository) Delete(id string) error {
	return r.DB.Delete(&entity.Article{}, id).Error
}
