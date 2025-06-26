package repository

import (
	"perpustakaan/contract"
	"perpustakaan/entity"
	"strings"

	"gorm.io/gorm"
)

type perpusRepo struct {
	db *gorm.DB
}

func implPerpusRepository(db *gorm.DB) contract.PerpusRepository {
	return &perpusRepo{
		db: db,
	}
}

func (r *perpusRepo) GetPerpus(id int) (*entity.Perpus, error) {
	var perpus entity.Perpus
	err := r.db.Table("perpus").Where("id = ?", id).First(&perpus).Error
	if err != nil {
		return nil, err
	}
	return &perpus, err
}

func (r *perpusRepo) CreatePerpus(perpus *entity.Perpus) error {
	return r.db.Table("perpus").Create(perpus).Error
}

func (r *perpusRepo) UpdatePerpus(id int, perpus *entity.Perpus) error {
	return r.db.Where("id = ?", id).Updates(&perpus).Error
}

func (r *perpusRepo) DeletePerpus(id int) error {
	return r.db.Where("id = ?", id).Delete(&entity.Perpus{}).Error
}

func (r *perpusRepo) UpdateStatus(id int, status string) error {
	return r.db.Model(&entity.Perpus{}).Where("id = ?", id).Update("status", status).Error
}

func (r *perpusRepo) SearchPerpusByJudul(judul string) ([]entity.Perpus, error) {
	var hasil []entity.Perpus
	err := r.db.Table("perpus").
		Where("LOWER(judul) LIKE ?", "%"+strings.ToLower(judul)+"%").
		Find(&hasil).Error

	return hasil, err
}

func (r *perpusRepo) CariJudul() ([]entity.Perpus, error) {
	var all []entity.Perpus
	err := r.db.Table("perpus").Find(&all).Error
	return all, err
}