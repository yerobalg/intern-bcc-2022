package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username     string `json:"username" gorm:"type:varchar(100);not null;unique"`
	Email        string `json:"email" gorm:"type:varchar(100);not null;unique"`
	Password     string `gorm:"type:text;not null"`
	Nama         string `json:"nama" gorm:"type:varchar(255);not null"`
	JenisKelamin bool   `json:"jenisKelamin" gorm:"type:boolean;not null"`
	RoleID       uint64 `gorm:"type:bigint;not null"`
	NomorHp      string `json:"nomorHP" gorm:"type:varchar(20);not null"`
	Role         Roles
}

type UserRegisterInput struct {
	Username     string `json:"username" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=6"`
	Nama         string `json:"nama" binding:"required"`
	NomorHp      string `json:"nomorHP" binding:"required"`
	RoleID 		   uint64 `json:"roleID" binding:"required"`
	JenisKelamin bool   `json:"jenisKelamin" binding:"required"`
}

type UserLoginInput struct {
	UsernameOrEmail string `json:"usernameOrEmail" binding:"required"`
	Password        string `json:"password" binding:"required,min=6"`
}
