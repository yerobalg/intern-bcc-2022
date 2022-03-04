package produk

import (
	"clean-arch-2/kategori"
	"clean-arch-2/user"
	"gorm.io/gorm"
)

type Produk struct {
	gorm.Model
	NamaProduk string     `json:"nama_produk" gorm:"type:varchar(100);not null"`
	Slug       string     `json:"slug" gorm:"type:varchar(100);not null;unique"`
	Harga      uint64     `json:"harga" gorm:"type:bigint;not null"`
	Stok       uint       `json:"stok" gorm:"type:int;not null"`
	Deskripsi  string     `json:"deskripsi" gorm:"type:text;not null"`
	IDSeller   uint64     `json:"idSeller" gorm:"type:bigint;not null;column:id_seller"`
	IsHiasan   bool       `json:"isHiasan" gorm:"type:boolean;not null"`
	Gender     string     `json:"gender" gorm:"type:varchar(10);"`
	Seller     user.Users `gorm:"foreignkey:IDSeller;references:ID;constraint:OnDelete:CASCADE"`
}

type Kategori_Produk struct {
	IDProduk   uint64            `json:"idProduk" gorm:"type:bigint;"`
	IDKategori uint64            `json:"idKategori" gorm:"type:bigint;"`
	Produk     Produk            `gorm:"foreignkey:IDProduk;references:ID;constraint:OnDelete:CASCADE"`
	Kategori   kategori.Kategori `gorm:"foreignkey:IDKategori;references:ID;constraint:OnDelete:CASCADE"`
}

type ProdukInput struct {
	NamaProduk string   `json:"namaProduk" binding:"required"`
	Harga      uint64   `json:"harga" binding:"required"`
	Stok       uint     `json:"stok" binding:"required"`
	Deskripsi  string   `json:"deskripsi" binding:"required"`
	Gender     string   `json:"gender"`
	IsHiasan   bool     `json:"isHiasan"`
	IDKategori uint64   `json:"idKategori"`
	IdTags     []uint64 `json:"idTags"`
}

func (Produk) TableName() string {
	return "produk"
}

func (Kategori_Produk) TableName() string {
	return "kategori_produk"
}
