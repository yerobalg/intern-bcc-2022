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

func (r *ProdukRepository) DeleteKategoriProduk(idProduk uint64) error {
	return r.Conn.Where("id_produk = ?", idProduk).Delete(&Kategori_Produk{}).Error
}

func (r *ProdukRepository) GetBySlug(slug string) (*Produk, error) {
	var produk Produk
	result := r.Conn.
		Preload("Seller").
		Preload("KategoriProduk").
		Preload("KategoriProduk.Kategori").
		First(&produk)
	return &produk, result.Error
}

func (r *ProdukRepository) GetByIdSeller(idSeller uint64) (*[]Produk, error) {
	var produk []Produk
	result := r.Conn.Where("id_seller = ?", idSeller).Find(&produk)
	return &produk, result.Error
}

func (r *ProdukRepository) Update(prod *Produk) error {
	return r.Conn.Raw(
		`
			UPDATE 
				"produk" 
			SET 
				"created_at" = ?, 
				"updated_at" = ?, 
				"deleted_at" = NULL, 
				"nama_produk" = ?, 
				"slug" = ?, 
				"harga" = ?, 
				"diskon" = ?, 
				"stok" = ?, 
				"deskripsi" = ?, 
				"id_seller" = ?, 
				"is_hiasan" = ?, 
				"gender" = ? 
			WHERE 
				"id" = ? 
				AND "produk"."deleted_at" IS NULL
		`,
		prod.CreatedAt,
		prod.UpdatedAt,
		prod.NamaProduk,
		prod.Slug,
		prod.Harga,
		prod.Diskon,
		prod.Stok,
		prod.Deskripsi,
		prod.IDSeller,
		prod.IsHiasan,
		prod.Gender,
		prod.ID,
	).Scan(&prod).Error
}

func (r *ProdukRepository) Delete(produk *Produk) error {
	return r.Conn.Delete(produk).Error
}
