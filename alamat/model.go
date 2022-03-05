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
	IDKabupaten   string           `json:"idKabupaten" gorm:"type:varchar(2);not null"`
	IDUser        uint64           `json:"idUser" gorm:"type:bigint;not null"`
	IsUser        bool             `gorm:"type:boolean;not null"`
	Kabupaten     daerah.Kabupaten `gorm:"foreignkey:IDKabupaten;references:IDKab"`
	User          user.Users       `gorm:"foreignkey:IDUser;references:ID;constraint:OnDelete:CASCADE"`
}

type AlamatInput struct {
	Nama          string `json:"nama" binding:"required"`
	AlamatLengkap string `json:"alamatLengkap" binding:"required"`
	KodePos       string `json:"kodePos" binding:"required,numeric,len=5"`
	IDKabupaten   string `json:"idKabupaten" binding:"required"`
}

func (Alamat) TableName() string {
	return "alamat"
}
