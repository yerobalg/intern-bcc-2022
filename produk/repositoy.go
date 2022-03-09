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

func (r *ProdukRepository) SaveGambarProduk(
	idProduk uint64,
	nama []string,
) error {
	var gambarProduk []Gambar_Produk
	for _, n := range nama {
		gambarProduk = append(gambarProduk, Gambar_Produk{
			IDProduk: idProduk,
			Nama:     n,
		})
	}
	return r.Conn.Create(&gambarProduk).Error
}

func (r *ProdukRepository) DeleteKategoriProduk(idProduk uint64) error {
	return r.Conn.Where("id_produk = ?", idProduk).Delete(&Kategori_Produk{}).Error
}

func (r *ProdukRepository) GetBySlug(slug string) (*Produk, error) {
	var produk Produk
	result := r.Conn.
		Preload("KategoriProduk").
		Preload("GambarProduk").
		Preload("KategoriProduk.Kategori").
		Where("slug = ?", slug).
		First(&produk)
	return &produk, result.Error
}

// func (r *ProdukRepository) GetByIdSeller(idSeller uint64) (*[]Produk, error) {
// 	var produk []Produk
// 	result := r.Conn.Where("id_seller = ?", idSeller).Find(&produk)
// 	return &produk, result.Error
// }

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
				"berat" = ?, 
				"diskon" = ?, 
				"stok" = ?, 
				"deskripsi" = ?, 
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
		prod.Berat,
		prod.Diskon,
		prod.Stok,
		prod.Deskripsi,
		prod.IsHiasan,
		prod.Gender,
		prod.ID,
	).Scan(&prod).Error
}

func (r *ProdukRepository) Delete(produk *Produk) error {
	return r.Conn.Delete(produk).Error
}

func (r *ProdukRepository) DeleteGambarProduk(
	idProduk uint64,
	nama string,
) error {
	gambar := &Gambar_Produk {
		IDProduk: idProduk,
		Nama: nama,
	}
	return r.Conn.Raw(
		`DELETE FEOM gambar_produk WHERE id_produk = ? AND nama = ?`,
		idProduk,
		nama,
	).Scan(&gambar).Error
}
