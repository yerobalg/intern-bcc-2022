package keranjang

import (
	"clean-arch-2/alamat"
	"os"
)

type KeranjangFormat struct {
	KeranjangProduk []KeranjangProduk `json:"keranjang"`
	TotalHarga      uint64            `json:"totalHarga"`
	TotalBerat      uint64            `json:"totalBerat"`
	Alamat          []alamat.GetAlamatFormatter   `json:"alamat"`
}
type KeranjangProduk struct {
	ID           uint64   `json:"id"`
	Slug         string   `json:"slug"`
	Diskon       float64  `json:"diskon"`
	Nama         string   `json:"nama"`
	Harga        uint64   `json:"harga"`
	JumlahBeli   uint     `json:"jumlahBeli"`
	Berat        uint     `json:"berat"`
	GambarProduk []string `json:"gambarProduk"`
}

func KeranjangFormatter(
	keranjang *[]Keranjang,
	alamatUser []alamat.GetAlamatFormatter,
) KeranjangFormat {
	var formattedKeranjang []KeranjangProduk
	totalHarga := 0.0
	totalBerat := uint64(0)

	for _, krj := range *keranjang {
		var gambar []string
		for _, gmb := range krj.Produk.GambarProduk {
			gambar = append(gambar, os.Getenv("BASE_URL")+"/"+gmb.Nama)
		}
		totalHarga += (1 - krj.Produk.Diskon) * float64(krj.Produk.Harga) * float64(krj.Jumlah)
		totalBerat += uint64(krj.Produk.Berat) * uint64(krj.Jumlah)
		formattedKeranjang = append(formattedKeranjang, KeranjangProduk{
			ID:           uint64(krj.ID),
			Slug:         krj.Produk.Slug,
			Diskon:       krj.Produk.Diskon,
			Nama:         krj.Produk.NamaProduk,
			Harga:        uint64(krj.Produk.Harga),
			JumlahBeli:   krj.Jumlah,
			Berat:        krj.Produk.Berat,
			GambarProduk: gambar,
		})
	}

	return KeranjangFormat{
		KeranjangProduk: formattedKeranjang,
		TotalHarga:      uint64(totalHarga),
		TotalBerat:      totalBerat,
		Alamat:          alamatUser,
	}
}
