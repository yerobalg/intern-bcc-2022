package produk

import (
	"clean-arch-2/kategori"
	"gorm.io/gorm"
)

type Produk struct {
	gorm.Model
	NamaProduk     string            `json:"nama_produk" gorm:"type:varchar(100);not null"`
	Slug           string            `json:"slug" gorm:"type:varchar(100);unique;notNull;size:256"`
	Harga          uint64            `json:"harga" gorm:"type:bigint;not null"`
	Diskon         float64           `json:"diskon" gorm:"type:double precision;default:0.0"`
	Stok           uint              `json:"stok" gorm:"type:int;not null"`
	Berat          uint              `json:"berat" gorm:"type:int;not null"`
	Deskripsi      string            `json:"deskripsi" gorm:"type:text;not null"`
	IsHiasan       bool              `json:"isHiasan" gorm:"type:boolean;not null"`
	Gender         string            `json:"gender" gorm:"type:varchar(10);"`
	KategoriProduk []Kategori_Produk `gorm:"foreignkey:IDProduk"`
}

type Gambar_Produk struct {
	IDProduk uint64 `json:"id_produk" gorm:"type:bigint;not null"`
	Nama     string `json:"nama" gorm:"type:varchar(100);not null"`
	Produk   Produk `gorm:"foreignkey:IDProduk;references:ID;constraint:OnDelete:CASCADE"`
}

func (Gambar_Produk) TableName() string {
	return "gambar_produk"
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
	Diskon     float64  `json:"diskon"`
	Gender     string   `json:"gender"`
	Berat      uint     `json:"berat"`
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
