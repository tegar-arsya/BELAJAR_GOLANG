package service


import (
    "portfolio-backend/internal/domain/entity"
    "portfolio-backend/internal/repository"
)

// AboutService interface untuk operasi CRUD pada About
type AboutService struct {
    Repo *repository.AboutRepository
}

// Create menyimpan data About baru ke database
func (s *AboutService) Create(data *entity.About) error {
    return s.Repo.Create(data)
}


// GetAll mengambil semua data About dari database
func (s *AboutService) GetAll() ([]entity.About, error) {
    return s.Repo.FindAll()
}

// update memperbarui data About berdasarkan ID
func (s *AboutService) Update(data *entity.About) error {
    return s.Repo.Update(data)
}
// Delete menghapus data About berdasarkan ID
func (s *AboutService) Delete(id string) error {
    return s.Repo.Delete(id)
}
