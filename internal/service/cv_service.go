package service

import (
    "portfolio-backend/internal/domain/entity"
    "portfolio-backend/internal/repository"
)
// CvService interface untuk operasi CRUD pada CV
type CvService struct {
    Repo *repository.CvRepository
}
// Create menyimpan data CV baru ke database
func (s *CvService) Create(data *entity.Cv) error {
    return s.Repo.Create(data)
}
// GetAll mengambil semua data CV dari database
func (s *CvService) GetAll() ([]entity.Cv, error) {
    return s.Repo.FindAll()
}
// Delete menghapus data CV berdasarkan ID
func (s *CvService) Delete(id string) error {
    return s.Repo.Delete(id)
}
