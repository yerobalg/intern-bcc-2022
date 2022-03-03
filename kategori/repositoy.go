package kategori

import (
	"gorm.io/gorm"
)

type KategoriRepository struct {
	Conn *gorm.DB
}

func NewKategoriRepository(Conn *gorm.DB) KategoriRepository {
	return KategoriRepository{Conn}
}

func (r *KategoriRepository) Save(kategori *Kategori) error {
	return r.Conn.Create(kategori).Error
}

func (r *KategoriRepository) Update(kategori *Kategori) error {
	return r.Conn.Save(kategori).Error
}

func (r *KategoriRepository) Delete(kategori *Kategori) error {
	return r.Conn.Delete(kategori).Error
}

func (r *KategoriRepository) GetById(id uint64) (Kategori, error) {
	var kategori Kategori
	result := r.Conn.Where("id = ?", id).First(&kategori)
	return kategori, result.Error
}

func (r *KategoriRepository) GetSemuaKategori() ([]Kategori, error) {
	var kategori []Kategori
	result := r.Conn.Find(&kategori)
	return kategori, result.Error
}
