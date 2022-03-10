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
		Where("is_paid = ?", false).
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

func (r *KeranjangRepository) GetSemuaMetodeBayar() ([]Metode_Pembayaran, error) {
	var metode_pembayaran []Metode_Pembayaran
	result := r.Conn.Find(&metode_pembayaran)
	return metode_pembayaran, result.Error
}

func (r *KeranjangRepository) GetMetodeByID(id uint64) (Metode_Pembayaran, error) {
	var metodePembayaran Metode_Pembayaran
	result := r.Conn.Where("id = ?", id).Find(&metodePembayaran)
	return metodePembayaran, result.Error
}

func (r *KeranjangRepository) KeranjangDipesan(idKeranjang []uint64) error {
	return r.Conn.Model(&Keranjang{}).
		Where("id IN ?", idKeranjang).
		Updates(Keranjang{IsPaid: true}).
		Error
}
