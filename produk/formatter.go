package produk

import (
	"clean-arch-2/kategori"
	"os"
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
	Gambar     []string                  `json:"gambar"`
	Kategori   kategori.KategoriFormat   `json:"kategori"`
	Tags       []kategori.KategoriFormat `json:"tags"`
	Hiasan     ProdukSearchFormat        `json:"hiasan"`
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
	hiasan ProdukSearchFormat,
) ProdukOutputFormat {
	daftarKategori := produk.KategoriProduk
	kategoriProduk := daftarKategori[len(daftarKategori)-1].Kategori
	var tagProduk []kategori.Kategori
	for i := 0; i < len(daftarKategori[:len(daftarKategori)-1]); i++ {
		tagProduk = append(tagProduk, daftarKategori[i].Kategori)
	}

	var daftarGambar []string
	for _, gambar := range produk.GambarProduk {
		daftarGambar = append(
			daftarGambar,
			os.Getenv("BASE_URL")+"/"+gambar.Nama,
		)
	}

	if len(daftarGambar) == 0 {
		daftarGambar = append(
			daftarGambar,
			os.Getenv("BASE_URL")+"/public/images/products/default.jpg",
		)
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
		Gambar:     daftarGambar,
		Kategori:   kategori.GetKategoriFormatter(&kategoriProduk),
		Tags:       kategori.GetSemuaKategoriFormatter(&tagProduk),
		Hiasan:     hiasan,
	}

}

type ProdukLiteFormat struct {
	Produk ProdukLite `json:"produk"`
	Gambar []string   `json:"gambar"`
}

type ProdukSearchFormat struct {
	JumlahHasil int64              `json:"jumlahHasil"`
	Produk      []ProdukLiteFormat `json:"daftarProduk"`
}

func ProdukSearchFormatter(
	hasilCari []ProdukLite,
	daftarGambar []Produk,
	jumlahHasil int64,
) ProdukSearchFormat {
	var daftarProduk []ProdukLiteFormat
	for _, produkCari := range hasilCari {
		for _, produk := range daftarGambar {
			if produkCari.ID == uint64(produk.ID) {
				var daftarGambar []string
				for _, gambar := range produk.GambarProduk {
					daftarGambar = append(
						daftarGambar,
						os.Getenv("BASE_URL")+"/"+gambar.Nama,
					)
				}
				if len(daftarGambar) == 0 {
					daftarGambar = append(
						daftarGambar,
						os.Getenv("BASE_URL")+"/public/images/products/default.jpg",
					)
				}
				daftarProduk = append(
					daftarProduk,
					ProdukLiteFormat{
						Produk: produkCari,
						Gambar: daftarGambar,
					},
				)
				break
			}
		}
	}

	return ProdukSearchFormat{
		JumlahHasil: jumlahHasil,
		Produk:      daftarProduk,
	}
}
