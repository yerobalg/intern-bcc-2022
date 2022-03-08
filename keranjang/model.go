package keranjang

import (
	"clean-arch-2/user"
	"gorm.io/gorm"
)

type InputKeranjang struct {
	IDSeller   uint64 `json:"idSeller"`
	Slug       string `json:"slug" binding:"required"`
	JumlahBeli uint   `json:"jumlahBeli"`
}

type Keranjang struct {
	gorm.Model
	IDUser uint64             `json:"idUser" gorm:"type:bigint;not null"`
	User   user.Users         `json:"user" gorm:"foreignkey:IDUser"`
	Produk []Keranjang_Produk `gorm:"foreignkey:IDKeranjang;"`
}

func (Keranjang) TableName() string {
	return "keranjang"
}

type Keranjang_Produk struct {
	IDProduk    uint64    `gorm:"type:bigint;not null"`
	Jumlah      uint      `gorm:"type:integer;not null"`
	IDKeranjang uint64    `gorm:"type:bigint;column:id_keranjang"`
	Keranjang   Keranjang `json:"keranjang" gorm:"foreignkey:IDKeranjang;references:ID;constraint:OnDelete:CASCADE"`
}

func (Keranjang_Produk) TableName() string {
	return "keranjang_produk"
}
