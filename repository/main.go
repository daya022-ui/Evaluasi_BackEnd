package repository

import (
	"perpustakaan/contract"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *contract.Repository {
	return &contract.Repository{
		// Code here
		// Example:
		// Example: implExampleRepository(db),
		Perpus: implPerpusRepository(db),
	}
}