package contract

import "perpustakaan/dto"

type Service struct {
	Perpus PerpusService
}

// type exampleService interface {
// Code here
// }

type PerpusService interface {
	GetPerpus(perpusID int) (*dto.PerpusResponse, error)
	CreatePerpus(payload *dto.PerpusRequest) (*dto.PerpusResponse, error)
	UpdatePerpus(id int, payload *dto.PerpusRequest) (*dto.PerpusResponse, error)
	DeletePerpus(id int) (*dto.PerpusResponse, error)
	PinjamBuku(id int) (*dto.PerpusResponse, error)
	KembalikanBuku(id int) error
	SearchPerpusByJudul(judul string) (*dto.PerpusListResponse, error)
	CariJudul() (*dto.PerpusListResponse, error)
}