package contract

import "perpustakaan/entity"

type Repository struct {
	Perpus PerpusRepository
}

// type exampleRepository interface {
// Code here
// }

type PerpusRepository interface {
	CreatePerpus(perpus *entity.Perpus) error
	GetPerpus(id int) (*entity.Perpus, error)
	UpdatePerpus(id int, perpus *entity.Perpus) error
	DeletePerpus(id int) error
	UpdateStatus(id int, status string) error
	SearchPerpusByJudul(judul string) ([]entity.Perpus, error)
	CariJudul() ([]entity.Perpus, error)
}