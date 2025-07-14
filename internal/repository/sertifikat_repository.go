package repository

import (
    "portfolio-backend/internal/domain/entity"
    "gorm.io/gorm"
)

type SertifikatRepository struct {
    DB *gorm.DB
}

func (r *SertifikatRepository) Create(s *entity.Sertifikat) error {
    return r.DB.Create(s).Error
}

func (r *SertifikatRepository) FindAll() ([]entity.Sertifikat, error) {
    var result []entity.Sertifikat
    err := r.DB.Find(&result).Error
    return result, err
}

func (r *SertifikatRepository) Delete(id string) error {
    return r.DB.Delete(&entity.Sertifikat{}, id).Error
}

func (r *SertifikatRepository) FindByID(id string) (*entity.Sertifikat, error) {
    var sertifikat entity.Sertifikat
    err := r.DB.First(&sertifikat, id).Error
    if err != nil {
        return nil, err
    }
    return &sertifikat, nil
}
