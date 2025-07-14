package repository

import (
    "portfolio-backend/internal/domain/entity"
    "gorm.io/gorm"
)

type PortofolioRepository struct {
    DB *gorm.DB
}
func (r *PortofolioRepository) Create(p *entity.Portfolio) error {
    return r.DB.Create(p).Error
}
func (r *PortofolioRepository) FindAll() ([]entity.Portfolio, error) {
    var result []entity.Portfolio
    err := r.DB.Find(&result).Error
    return result, err
}
func (r *PortofolioRepository) FindByID(id string) (*entity.Portfolio, error) {
    var portofolio entity.Portfolio
    err := r.DB.First(&portofolio, "id = ?", id).Error
    if err != nil {
        return nil, err
    }
    return &portofolio, nil
}
func (r *PortofolioRepository) Update(p *entity.Portfolio) error {
    return r.DB.Save(p).Error
}
func (r *PortofolioRepository) Delete(id string) error {
    return r.DB.Delete(&entity.Portfolio{}, id).Error
}
