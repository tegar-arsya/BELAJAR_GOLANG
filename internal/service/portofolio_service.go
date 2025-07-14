package service

import (
    "portfolio-backend/internal/domain/entity"
    "portfolio-backend/internal/repository"
)

type PortofolioService struct {
    Repo *repository.PortofolioRepository
}

func (s *PortofolioService) Create(data *entity.Portfolio) error {
    return s.Repo.Create(data)
}


func (s *PortofolioService) GetAll() ([]entity.Portfolio, error) {
    return s.Repo.FindAll()
}

func (s *PortofolioService) GetByID(id string) (*entity.Portfolio, error) {
    return s.Repo.FindByID(id)
}

func (s *PortofolioService) Update(data *entity.Portfolio) error {
    return s.Repo.Update(data)
}
func (s *PortofolioService) Delete(id string) error {
    return s.Repo.Delete(id)
}


