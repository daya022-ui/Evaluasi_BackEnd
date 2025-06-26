package entity

type Perpus struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement;not null;<-create"`
	Judul     string `gorm:"column:judul;type:varchar(255);not null"`
	Penulis   string `gorm:"column:penulis;type:varchar(255);not null"`
	Status    string `gorm:"column:status;type:varchar(10);not null;default:'available'"`
	UpdatedAt string `gorm:"column:updated_at;type:timestamp;not null;default:now()"`
	CreatedAt string `gorm:"column:created_at;type:timestamp;not null;default:now()"`
}

func (e *Perpus) TableName() string {
	return "perpus"
}