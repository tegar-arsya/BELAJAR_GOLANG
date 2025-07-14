package service

import (
	"portfolio-backend/internal/domain/entity"
	"portfolio-backend/internal/repository"
)

type ArticleService struct {
	Repo *repository.ArticleRepository
}

func (s *ArticleService) Create(data *entity.Article) error {
	return s.Repo.Create(data)
}

func (s *ArticleService) GetAll() ([]entity.Article, error) {
	return s.Repo.FindAll()
}

func (s *ArticleService) GetByID(id string) (*entity.Article, error) {
	return s.Repo.FindByID(id)
}

func (s *ArticleService) Update(data *entity.Article) error {
	return s.Repo.Update(data)
}

func (s *ArticleService) Delete(id string) error {
	return s.Repo.Delete(id)
}
