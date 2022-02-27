package daerah

import (
	"fmt"
  "gorm.io/gorm"
)

type DaerahRepository struct {
	Conn *gorm.DB
}

func NewDaerahRepository(Conn *gorm.DB) DaerahRepository {
	return DaerahRepository{Conn}
}

func (r *DaerahRepository) GetDaerah(
	daerah *[]OutputDaerah,
	id string,
	kolomAsal string,
	kolomTujuan string,
	tabelTujuan string,
) error {
	return r.Conn.Raw(
		fmt.Sprintf(
			"SELECT %s AS id, nama FROM %s WHERE %s = ?",
			kolomAsal,
			tabelTujuan,
			kolomTujuan,
		),
		id,
	).Find(&daerah).Error
}

func (r *DaerahRepository) GetSemuaProvinsi(provinsi *[]Provinsi) error {
  return r.Conn.Find(&provinsi).Error
}
