package keranjang

import (
	"gorm.io/gorm"
)

type KeranjangRepository struct {
	Conn *gorm.DB
}

func NewKeranjangRepository(Conn *gorm.DB) KeranjangRepository {
	return KeranjangRepository{Conn}
}

func (r *KeranjangRepository) AddKeranjang(keranjang *Keranjang) error {
	return r.Conn.Create(keranjang).Error
}

func (r *KeranjangRepository) GetKeranjangUser(
	idUser uint64,
) ([]Keranjang, error) {
	var keranjang []Keranjang
	result := r.Conn.
		Where("id_user = ?", idUser).
		Preload("Produk").
		Preload("Produk.GambarProduk").
		Find(&keranjang)
	return keranjang, result.Error
}

func (r *KeranjangRepository) GetKeranjangBatch(
	idBatch []uint64,
) ([]Keranjang, error) {
	var keranjang []Keranjang
	res := r.Conn.
		Preload("Produk").
		Preload("Produk.GambarProduk").
		Find(&keranjang, idBatch)
	return keranjang, res.Error
}
