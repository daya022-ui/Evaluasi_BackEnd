package service

import (
	"errors"
	"fmt"
	"perpustakaan/contract"
	"perpustakaan/dto"
	"perpustakaan/entity"
	"net/http"
	"strings"
)

type PerpusService struct {
	perpusRepository contract.PerpusRepository
}

func implPerpusService(repo *contract.Repository) contract.PerpusService {
	return &PerpusService{
		perpusRepository: repo.Perpus,
	}
}

func (s *PerpusService) GetPerpus(perpusID int) (*dto.PerpusResponse, error) {
	perpus, err := s.perpusRepository.GetPerpus(perpusID)
	if err != nil {
		return nil, err
	}

	response := &dto.PerpusResponse{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mendapatkan data",
		Data: dto.PerpusData{
			ID:      perpus.ID,
			Judul:   perpus.Judul,
			Penulis: perpus.Penulis,
			Status:  perpus.Status,
		},
	}
	return response, nil
}

func (s *PerpusService) CreatePerpus(payload *dto.PerpusRequest) (*dto.PerpusResponse, error) {
	perpus := &entity.Perpus{
		Judul:   payload.Judul,
		Penulis: payload.Penulis,
		Status:  payload.Status,
	}

	err := s.perpusRepository.CreatePerpus(perpus)
	if err != nil {
		return nil, err
	}

	response := &dto.PerpusResponse{
		StatusCode: http.StatusCreated,
		Message:    "Berhasil membuat data",
		Data: dto.PerpusData{
			ID:      perpus.ID,
			Judul:   perpus.Judul,
			Penulis: perpus.Penulis,
			Status:  perpus.Status,
		},
	}

	return response, nil
}

func (s *PerpusService) UpdatePerpus(id int, payload *dto.PerpusRequest) (*dto.PerpusResponse, error) {
	perpus := &entity.Perpus{
		Judul:   payload.Judul,
		Penulis: payload.Penulis,
		Status:  payload.Status,
	}

	err := s.perpusRepository.UpdatePerpus(id, perpus)
	if err != nil {
		return nil, err
	}

	response := &dto.PerpusResponse{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mengubah data",
		Data: dto.PerpusData{
			ID:      perpus.ID,
			Judul:   perpus.Judul,
			Penulis: perpus.Penulis,
			Status:  perpus.Status,
		},
	}

	return response, nil
}

func (s *PerpusService) DeletePerpus(id int) (*dto.PerpusResponse, error) {
	err := s.perpusRepository.DeletePerpus(id)
	if err != nil {
		return nil, err
	}

	response := &dto.PerpusResponse{
		StatusCode: http.StatusOK,
		Message:    "Berhasil menghapus data",
	}

	return response, nil
}

func (s *PerpusService) PinjamBuku(id int) (*dto.PerpusResponse, error) {
	perpus, err := s.perpusRepository.GetPerpus(id)
	if err != nil {
		return nil, err
	}

	if perpus.Status != "available" {
		return nil, errors.New("buku sedang dipinjam")
	}

	err = s.perpusRepository.UpdateStatus(id, "borrowed")
	if err != nil {
		return nil, err
	}

	filename := strings.ToLower(strings.ReplaceAll(perpus.Judul, " ", "_")) + ".pdf"
	downloadLink := fmt.Sprintf("http://localhost:8080/files/%s", filename)

	response := &dto.PerpusResponse{
		StatusCode: http.StatusOK,
		Message:    "Buku berhasil dipinjam",
		Data: dto.PerpusData{
			ID:           perpus.ID,
			Judul:        perpus.Judul,
			Penulis:      perpus.Penulis,
			DownloadLink: downloadLink,
			Status:       "borrowed",
		},
	}
	return response, nil
}

func (s *PerpusService) KembalikanBuku(id int) error {
	return s.perpusRepository.UpdateStatus(id, "available")
}

func (s *PerpusService) SearchPerpusByJudul(judul string) (*dto.PerpusListResponse, error) {
	hasil, err := s.perpusRepository.SearchPerpusByJudul(judul)
	if err != nil {
		return nil, err
	}

	var list []dto.PerpusData
	for _, p := range hasil {
		list = append(list, dto.PerpusData{
			ID:      p.ID,
			Judul:   p.Judul,
			Penulis: p.Penulis,
			Status:  p.Status,
		})
	}

	return &dto.PerpusListResponse{
		StatusCode: http.StatusOK,
		Message:    "Hasil pencarian judul",
		Data:       list,
	}, nil
}

func (s *PerpusService) CariJudul() (*dto.PerpusListResponse, error) {
	data, err := s.perpusRepository.CariJudul()
	if err != nil {
		return nil, err
	}

	var list []dto.PerpusData
	for _, p := range data {
		list = append(list, dto.PerpusData{
			ID:      p.ID,
			Judul:   p.Judul,
			Penulis: p.Penulis,
			Status:  p.Status,
		})
	}

	return &dto.PerpusListResponse{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mengambil seluruh data",
		Data:       list,
	}, nil
}