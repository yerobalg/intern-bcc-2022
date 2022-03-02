package alamat

import (
	"gorm.io/gorm"
)

type AlamatRepository struct {
	Conn *gorm.DB
}

func NewAlamatRepository(Conn *gorm.DB) AlamatRepository {
	return AlamatRepository{Conn}
}

func (r *AlamatRepository) Save(alamat *Alamat) (error) {
	return r.Conn.Create(&alamat).Error
}