package kategori

import (
	"gorm.io/gorm"
)

type Kategori struct {
	gorm.Model
	Nama  string `json:"nama" gorm:"type:varchar(100);not null"`
	IsTag bool   `gorm:"type:boolean;default:false"`
}

type KategoriInput struct {
	Nama  string `json:"nama" binding:"required"`
	IsTag bool   `json:"isTag"`
}

func (Kategori) TableName() string {
	return "kategori"
}
