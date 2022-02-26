package alamat

import (
	"clean-arch-2/daerah"
	"clean-arch-2/user"
	"gorm.io/gorm"
)

type Alamat struct {
	gorm.Model
	Nama          string           `json:"nama" gorm:"type:varchar(50);not null"`
	AlamatLengkap string           `json:"alamatLengkap" gorm:"type:text;not null"`
	KodePos       string           `json:"kodePos" gorm:"type:varchar(10);not null"`
	IDKelurahan   string           `json:"idKelurahan" gorm:"type:varchar(10);not null"`
	IDUser        uint64           `json:"idUser" gorm:"type:bigint;not null"`
	User          user.Users       `json:"user" gorm:"foreignkey:IDUser;references:ID"`
	Kelurahan     daerah.Kelurahan `json:"kelurahan" gorm:"foreignkey:IDKelurahan;references:IDKel"`
}

func (Alamat) TableName() string {
	return "alamat"
}
