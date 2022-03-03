package produk

import (
	"gorm.io/gorm"
)

type Produk struct {
	gorm.Model
	NamaProduk string `json:"nama_produk" gorm:"type:varchar(100);not null"`
	Slug       string `json:"slug" gorm:"type:varchar(100);not null;unique"`
	Harga      uint64 `json:"harga" gorm:"type:bigint;not null"`
	Stok       uint   `json:"stok" gorm:"type:int;not null"`
	Deskripsi  string `json:"deskripsi" gorm:"type:text;not null"`
	IDSeller   uint64 `json:"idSeller" gorm:"type:bigint;not null;column:id_seller"`
	IDKategori uint64 `json:"idKategori" gorm:"type:bigint;not null;column:id_kategori"`
}
