package keranjang

import (
	"clean-arch-2/produk"
	"clean-arch-2/user"
	"gorm.io/gorm"
)

type InputKeranjang struct {
	Slug       string `json:"slug" binding:"required"`
	JumlahBeli uint   `json:"jumlahBeli"`
}

type InputKonfirmasi struct {
	IDAlamat    uint64   `json:"idAlamat"`
	IDKeranjang []uint64 `json:"idKeranjang"`
}

type DaftarBeliProduk struct {
	Slug       string `json:"slug" binding:"required"`
	JumlahBeli uint   `json:"jumlahBeli"`
}

type Keranjang struct {
	gorm.Model
	IDUser   uint64        `json:"idUser" gorm:"type:bigint;not null"`
	IDProduk uint64        `gorm:"type:bigint;not null"`
	Jumlah   uint          `gorm:"type:integer;not null"`
	User     user.Users    `json:"user" gorm:"foreignkey:IDUser;constraint:OnDelete:CASCADE"`
	Produk   produk.Produk `json:"produk" gorm:"foreignkey:IDProduk;constraint:OnDelete:CASCADE"`
}

func (Keranjang) TableName() string {
	return "keranjang"
}
type Metode_Pembayaran struct {
	ID          uint64 `json:"id" gorm:"primary_key;auto_increment:true"`
	Jenis       string `json:"jenis" gorm:"type:varchar(255);not null"`
	Nama        string `json:"nama" gorm:"type:varchar(255);not null"`
	BiayaAdmin  uint64 `json:"biayaAdmin" gorm:"type:bigint;not null"`
	ExpiredTime uint64 `json:"expiredTime" gorm:"type:bigint;not null"`
}

func (Metode_Pembayaran) TableName() string {
	return "metode_pembayaran"
}
