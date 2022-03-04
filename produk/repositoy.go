package produk

import (
	"gorm.io/gorm"
)

type ProdukRepository struct {
	Conn *gorm.DB
}

func NewProdukRepository(Conn *gorm.DB) ProdukRepository {
	return ProdukRepository{Conn}
}

func (r *ProdukRepository) Save(produk *Produk) error {
	return r.Conn.Create(produk).Error
}

func (r *ProdukRepository) SaveKategoriProduk(
	idKategori []uint64,
	idProduk uint64,
) error {
	var kategoriProduk []Kategori_Produk
	for _, id := range idKategori {
		kategoriProduk = append(kategoriProduk, Kategori_Produk{
			IDKategori: id,
			IDProduk:   idProduk,
		})
	}
	return r.Conn.Create(&kategoriProduk).Error
}

func (r *ProdukRepository) GetBySlug(slug string) (*Produk, error) {
	var produk Produk
	result := r.Conn.Where("slug = ?", slug).First(&produk)
	return &produk, result.Error
}

func (r *ProdukRepository) GetByIdSeller(idSeller uint64) (*[]Produk, error) {
	var produk []Produk
	result := r.Conn.Where("id_seller = ?", idSeller).Find(&produk)
	return &produk, result.Error
}

func (r *ProdukRepository) Update(produk *Produk) error {
	return r.Conn.Save(produk).Error
}

func (r *ProdukRepository) Delete(produk *Produk) error {
	return r.Conn.Delete(produk).Error
}
