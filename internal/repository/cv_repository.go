package repository

import (
    "portfolio-backend/internal/domain/entity"
    "gorm.io/gorm"
)

// CvRepository interface untuk operasi CRUD pada CV
type CvRepository struct {
    DB *gorm.DB
}

// Create menyimpan data CV baru ke database
func (r *CvRepository) Create(cv *entity.Cv) error {
    return r.DB.Create(cv).Error
}

// FindAll mengambil semua data CV dari database
func (r *CvRepository) FindAll() ([]entity.Cv, error) {
    var result []entity.Cv
    err := r.DB.Find(&result).Error
    return result, err
}

// Delete menghapus data CV berdasarkan ID
func (r *CvRepository) Delete(id string) error {
    return r.DB.Delete(&entity.Cv{}, id).Error
}
