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
	KodePos       string           `json:"kodePos" gorm:"type:varchar(5);not null;"`
	IDKelurahan   string           `json:"idKelurahan" gorm:"type:varchar(10);not null"`
	IDUser        uint64           `json:"idUser" gorm:"type:bigint;not null"`
	IsUser        bool             `gorm:"type:boolean;not null"`
	User          user.Users       `json:"user" gorm:"foreignkey:IDUser;references:ID;constraint:OnDelete:CASCADE"`
	Kelurahan     daerah.Kelurahan `json:"kelurahan" gorm:"foreignkey:IDKelurahan;references:IDKel"`
}

type AlamatInput struct {
	Nama          string `json:"nama" binding:"required"`
	AlamatLengkap string `json:"alamatLengkap" binding:"required"`
	KodePos       string `json:"kodePos" binding:"required,numeric,len=5"`
	IDKelurahan   string `json:"idKelurahan" binding:"required"`
}

func (Alamat) TableName() string {
	return "alamat"
}
