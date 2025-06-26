package dto

type PerpusRequest struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Judul     string `json:"judul"`
	Penulis   string `json:"penulis"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

type PerpusData struct {
	ID           int    `json:"id"`
	Judul        string `json:"judul"`
	Penulis      string `json:"penulis"`
	Status       string `json:"status"`
	DownloadLink string `json:"download_link"`
}

type PerpusResponse struct {
	StatusCode int        `json:"status_code"`
	Message    string     `json:"message"`
	Data       PerpusData `json:"data"`
}

type PerpusListResponse struct {
	StatusCode int          `json:"status_code"`
	Message    string       `json:"message"`
	Data       []PerpusData `json:"data"`
}

// Error implements error.
func (i *PerpusResponse) Error() string {
	panic("unimplemented")
}