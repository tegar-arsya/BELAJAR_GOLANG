package repository
import (
    "portfolio-backend/internal/domain/entity"
    "gorm.io/gorm"
)

// AboutRepository interface untuk operasi CRUD pada About
type AboutRepository struct {
    DB *gorm.DB
}

// Create menyimpan data About baru ke database
func (r *AboutRepository) Create(a *entity.About) error {
    return r.DB.Create(a).Error
}
// FindAll mengambil semua data About dari database
func (r *AboutRepository) FindAll() ([]entity.About, error) {
    var result []entity.About
    err := r.DB.Find(&result).Error
    return result, err
}

// Delete menghapus data About berdasarkan ID
func (r *AboutRepository) Delete(id string) error {
    return r.DB.Delete(&entity.About{}, id).Error
}

// Update memperbarui data About berdasarkan ID
func (r *AboutRepository) Update(a *entity.About) error {
    return r.DB.Save(a).Error
}
