package user

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
	NomorHp      string `json:"nomorHP" gorm:"type:varchar(20);not null;unique"`
	Role         Roles
}

type UserRegisterInput struct {
	Username     string `json:"username" binding:"required,min=3,max=15"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=6,max=20"`
	Nama         string `json:"nama" binding:"required"`
	NomorHp      string `json:"nomorHP" binding:"required,numeric"`
	RoleID 		   uint64 `json:"roleID" binding:"required"`
	JenisKelamin bool   `json:"jenisKelamin" binding:"required"`
}

type UserLoginInput struct {
	UsernameOrEmail string `json:"usernameOrEmail" binding:"required"`
	Password        string `json:"password" binding:"required,min=6"`
}

type Roles struct {
	ID 	 uint   `json:"id" gorm:"primary_key"`
	Nama string `json:"name" gorm:"type:varchar(100);not null"`
}