package produk

import (
	"fmt"
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

func (r *ProdukRepository) GetHiasanProduk(idKategori uint64) ([]ProdukLite, error, int64) {
	var prod []ProdukLite
	res := r.Conn.
		Model(&Produk{}).
		Select(`
			produk.id,
			produk.slug, 
			produk.nama_produk, 
			produk.harga, 
			produk.diskon
		`).
		Joins("LEFT JOIN kategori_produk ON produk.id = kategori_produk.id_produk").
		Where("kategori_produk.id_kategori = ?", idKategori).
		Where("produk.is_hiasan = ?", true).
		Find(&prod)
	return prod, res.Error, res.RowsAffected
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
) error {
	return r.Conn.Delete(&Gambar_Produk{}, "id_produk = ?", idProduk).Error
}

func (r *ProdukRepository) CariProdukTerbaru(
	idKategori []uint64,
	kataKunci string,
	maksHarga uint64,
	gender string,
	page int,
) ([]ProdukLite, error, int64) {
	offset := (page - 1) * 12
	var prod []ProdukLite
	trx := r.Conn.Model(&Produk{}).
		Select(`
			produk.id,
			produk.slug, 
			produk.nama_produk, 
			produk.harga, 
			produk.diskon
		`)

	if len(idKategori) > 0 {
		trx = trx.
			Joins("LEFT JOIN kategori_produk ON produk.id = kategori_produk.id_produk").
			Where("kategori_produk.id_kategori IN (?)", idKategori)
	}
	if kataKunci != "" {
		trx = trx.Where("produk.nama_produk ILIKE ?", "%"+kataKunci+"%")
	}
	if maksHarga != 0 {
		trx = trx.Where("(1-produk.diskon)*produk.harga <= ?", maksHarga)
	}
	if gender != "BOTH" {
		trx = trx.Where("produk.gender = ?", gender)
	}

	res := trx.
		Where("produk.is_hiasan = ?", false).
		Order("produk.created_at DESC").
		Limit(12).Offset(offset).Find(&prod)
	return prod, res.Error, res.RowsAffected
}

func (r *ProdukRepository) CariProdukTerlaris(
	idKategori []uint64,
	kataKunci string,
	maksHarga uint64,
	gender string,
	page int,
) ([]ProdukLite, error, int64) {
	offset := (page - 1) * 12
	var prod []ProdukLite
	trx := r.Conn.Model(&Produk{}).
		Select(`
			count(keranjang.id), 
			produk.id,
			produk.slug, 
			produk.nama_produk, 
			produk.harga, 
			produk.diskon
		`).
		Joins("LEFT JOIN keranjang ON produk.id = keranjang.id_produk AND keranjang.is_paid = true")

	if len(idKategori) > 0 {
		trx = trx.
			Joins("LEFT JOIN kategori_produk ON produk.id = kategori_produk.id_produk").
			Where("kategori_produk.id_kategori IN (?)", idKategori)
	}
	if kataKunci != "" {
		trx = trx.Where("produk.nama_produk ILIKE ?", "%"+kataKunci+"%")
	}
	if maksHarga != 0 {
		trx = trx.Where("(1-produk.diskon)*produk.harga <= ?", maksHarga)
	}
	if gender != "BOTH" {
		trx = trx.Where("produk.gender = ?", gender)
	}

	res := trx.
		Where("produk.is_hiasan = ?", false).
		Group("produk.id, produk.nama_produk, produk.harga, produk.diskon").
		Order("count(keranjang.id) DESC, harga").
		Limit(12).Offset(offset).Find(&prod)
	return prod, res.Error, res.RowsAffected
}

func (r *ProdukRepository) GetGambarProduk(
	idProduk []uint64,
) ([]Produk, error) {
	var prod []Produk
	res := r.Conn.
		Preload("GambarProduk").
		Where("id IN (?)", idProduk).
		Find(&prod)
	return prod, res.Error
}

func (r *ProdukRepository) KurangiStok(
	idProduk []uint64,
	jumlah []uint,
) error {
	return r.Conn.Transaction(func(tx *gorm.DB) error {
		for i, id := range idProduk {
			if err := tx.Model(&Produk{}).
				Where("id = ?", id).
				Update("stok", gorm.Expr("stok - ?", jumlah[i])).
				Error; err != nil {
				return err
			}
			fmt.Println("tes")
		}
		return nil
	})
}
