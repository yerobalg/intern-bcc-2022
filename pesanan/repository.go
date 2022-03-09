package pesanan

import (
	"gorm.io/gorm"
)

type PesananRepository struct {
	Conn *gorm.DB
}

func NewPesananRepository(Conn *gorm.DB) PesananRepository {
	return PesananRepository{Conn}
}

func (r *PesananRepository) AddPesanan(pesanan *Pesanan) error {
	return r.Conn.Create(pesanan).Error
}

func (r *PesananRepository) AddKeranjangPesanan(
	idPesanan uint64,
	idKeranjang []uint64,
) error {
	var keranjangPesanan []Keranjang_Pesanan
	for _, id := range idKeranjang {
		keranjangPesanan = append(keranjangPesanan, Keranjang_Pesanan{
			IDPesanan:   idPesanan,
			IDKeranjang: id,
		})
	}

	return r.Conn.Create(&keranjangPesanan).Error
}

func (r *PesananRepository) GetPesanan(idPesanan uint64) (Pesanan, error) {
	var pesanan Pesanan
	err := r.Conn.
		Where("id = ?", idPesanan).
		Preload("Keranjang_Pesanan").
		Find(&pesanan).Error
	return pesanan, err
}

func (r *PesananRepository) DeletePesanan(pesanan *Pesanan) error {
	return r.Conn.Delete(pesanan).Error
}

func (r *PesananRepository) DeleteKeranjangPesanan(
	idPesanan uint64,
) error {
	return r.Conn.
		Where("id_pesanan = ?", idPesanan).
		Delete(&Keranjang_Pesanan{}).
		Error
}
