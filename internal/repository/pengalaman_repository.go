package repository

import (
    "portfolio-backend/internal/domain/entity"
    "gorm.io/gorm"
)

// PengalamanRepository interface untuk operasi CRUD pada pengalaman
type PengalamanRepository struct {
    DB *gorm.DB
}

// Create menyimpan data pengalaman baru ke database
func (r *PengalamanRepository) Create(p *entity.Pengalaman) error {
    return r.DB.Create(p).Error
}

// FindAll mengambil semua data pengalaman dari database
func (r *PengalamanRepository) FindAll() ([]entity.Pengalaman, error) {
    var result []entity.Pengalaman
    err := r.DB.Find(&result).Error
    return result, err
}

// Delete menghapus data pengalaman berdasarkan ID
func (r *PengalamanRepository) Delete(id string) error {
    return r.DB.Delete(&entity.Pengalaman{}, id).Error
}

// FindByID mengambil data pengalaman berdasarkan ID
func (r *PengalamanRepository) FindByID(id string) (*entity.Pengalaman, error) {
    var pengalaman entity.Pengalaman
    err := r.DB.First(&pengalaman, id).Error
    if err != nil {
        return nil, err
    }
    return &pengalaman, nil
}
