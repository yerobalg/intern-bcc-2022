package daerah

import (
	"gorm.io/gorm"
)

type DaerahRepository struct {
	Conn *gorm.DB
}

func NewDaerahRepository(Conn *gorm.DB) DaerahRepository {
	return DaerahRepository{Conn}
}

func (r *DaerahRepository) GetKabupaten(idProvinsi string) ([]Kabupaten, error) {
	var kabupaten []Kabupaten
	res := r.Conn.Where("id_prov = ?", idProvinsi).Find(&kabupaten)
	return kabupaten, res.Error
}

func (r *DaerahRepository) GetSemuaProvinsi(provinsi *[]Provinsi) error {
	return r.Conn.Find(&provinsi).Error
}
