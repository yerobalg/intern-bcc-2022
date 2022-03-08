package produk

import (
	"clean-arch-2/kategori"
)

type ProdukInputFormat struct {
	NamaProduk string   `json:"nama"`
	Slug       string   `json:"slug"`
	Harga      uint64   `json:"harga"`
	Diskon     float64  `json:"diskon"`
	Stok       uint     `json:"stok"`
	Deskripsi  string   `json:"deskripsi"`
	Gender     string   `json:"gender"`
	Berat      uint     `json:"berat"`
	IsHiasan   bool     `json:"isHiasan"`
	IdKategori uint64   `json:"idKategori"`
	IdTags     []uint64 `json:"idTags"`
}
type ProdukOutputFormat struct {
	NamaProduk string                    `json:"nama"`
	Slug       string                    `json:"slug"`
	Harga      uint64                    `json:"harga"`
	Diskon     float64                   `json:"diskon"`
	Stok       uint                      `json:"stok"`
	Deskripsi  string                    `json:"deskripsi"`
	Gender     string                    `json:"gender"`
	Berat      uint                      `json:"berat"`
	IsHiasan   bool                      `json:"isHiasan"`
	Kategori   kategori.KategoriFormat   `json:"kategori"`
	Tags       []kategori.KategoriFormat `json:"tags"`
}

type Seller struct {
	ID      uint64 `json:"id"`
	Nama    string `json:"nama"`
	NomorHp string `json:"nomorHP"`
}

func ProdukInputFormatter(
	produk *Produk,
	idKategori uint64,
	IdTags []uint64,
) ProdukInputFormat {
	return ProdukInputFormat{
		NamaProduk: produk.NamaProduk,
		Slug:       produk.Slug,
		Harga:      produk.Harga,
		Diskon:     produk.Diskon,
		Stok:       produk.Stok,
		Deskripsi:  produk.Deskripsi,
		Gender:     produk.Gender,
		Berat:      produk.Berat,
		IsHiasan:   produk.IsHiasan,
		IdKategori: idKategori,
		IdTags:     IdTags,
	}
}

func ProdukOutputFormatter(
	produk Produk,
) ProdukOutputFormat {
	daftarKategori := produk.KategoriProduk
	kategoriProduk := daftarKategori[len(daftarKategori)-1].Kategori
	var tagProduk []kategori.Kategori

	for i := 0; i < len(daftarKategori[:len(daftarKategori)-1]); i++ {
		tagProduk = append(tagProduk, daftarKategori[i].Kategori)
	}

	return ProdukOutputFormat{
		NamaProduk: produk.NamaProduk,
		Slug:       produk.Slug,
		Harga:      produk.Harga,
		Diskon:     produk.Diskon,
		Stok:       produk.Stok,
		Deskripsi:  produk.Deskripsi,
		Gender:     produk.Gender,
		Berat:      produk.Berat,
		IsHiasan:   produk.IsHiasan,
		Kategori:   kategori.GetKategoriFormatter(&kategoriProduk),
		Tags:       kategori.GetSemuaKategoriFormatter(&tagProduk),
	}
}
