package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username     string `json:"username" gorm:"type:varchar(100);not null"`
	Email        string `json:"email" gorm:"type:varchar(100);not null"`
	Password     string `json:"password" gorm:"type:text;not null"`
	Nama         string `json:"name" gorm:"type:varchar(255);not null"`
	JenisKelamin bool   `json:"isMale" gorm:"type:boolean;not null"`
	TanggalLahir string `json:"birthday" gorm:"type:date;not null"`
	Alamat       string `json:"address" gorm:"type:varchar(255);not null"`
	RoleID       uint64 `gorm:"type:bigint;not null"`
	Role         Roles  `json:"role"`
}
