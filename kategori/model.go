package kategori

import "gorm.io/gorm"

type Kategori struct {
	gorm.Model
	Nama string `json:"nama" gorm:"type:varchar(100);not null"`
}

type KategoriInput struct {
	Nama string `json:"nama" binding:"required"`
}

func (Kategori) TableName() string {
	return "kategori"
}
