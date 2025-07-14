package service

import (
    "portfolio-backend/internal/domain/entity"
    "portfolio-backend/internal/repository"
)

type PengalamanService struct {
    Repo *repository.PengalamanRepository
}
func (s *PengalamanService) Create(data *entity.Pengalaman) error {
    return s.Repo.Create(data)
}

func (s *PengalamanService) GetAll() ([]entity.Pengalaman, error) {
    return s.Repo.FindAll()
}
func (s *PengalamanService) Delete(id string) error {
    return s.Repo.Delete(id)
}

func (s *PengalamanService) GetByID(id string) (*entity.Pengalaman, error) {
    return s.Repo.FindByID(id)
}
