package service

import (
    "portfolio-backend/internal/domain/entity"
    "portfolio-backend/internal/repository"
)

type SertifikatService struct {
    Repo *repository.SertifikatRepository
}

func (s *SertifikatService) Create(data *entity.Sertifikat) error {
    return s.Repo.Create(data)
}

func (s *SertifikatService) GetAll() ([]entity.Sertifikat, error) {
    return s.Repo.FindAll()
}

func (s *SertifikatService) Delete(id string) error {
    return s.Repo.Delete(id)
}

func (s *SertifikatService) GetByID(id string) (*entity.Sertifikat, error) {
    return s.Repo.FindByID(id)
}

func (s *SertifikatService) Update(data *entity.Sertifikat) error {
    // Sertifikat does not have an update method in the repository, so we can return an error or implement it if needed
    return s.Repo.DB.Save(data).Error
}
