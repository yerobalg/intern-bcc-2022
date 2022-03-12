package keranjang

import (
	"clean-arch-2/alamat"
	"clean-arch-2/user"
	"clean-arch-2/utilities"
)

type KeranjangFormat struct {
	KeranjangProduk []KeranjangProduk           `json:"keranjang"`
	TotalHarga      uint64                      `json:"totalHarga"`
	TotalBerat      uint64                      `json:"totalBerat"`
	Alamat          []alamat.GetAlamatFormatter `json:"alamat"`
}

type KonfirmasiPesananFormat struct {
	User             user.ProfilUserFormat     `json:"user"`
	AlamatUser       alamat.GetAlamatFormatter `json:"alamatUser"`
	AlamatAdmin      alamat.GetAlamatFormatter `json:"alamatAdmin"`
	KeranjangProduk  []KeranjangProduk         `json:"produk"`
	Pengiriman       []utilities.OngkirFormat  `json:"pengiriman"`
	MetodePembayaran []Metode_Pembayaran       `json:"metodePembayaran"`
	TotalHarga       uint64                    `json:"totalHarga"`
	TotalBerat       uint64                    `json:"totalBerat"`
}

type Users struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Nama     string `json:"nama"`
}
type KeranjangProduk struct {
	ID           uint64   `json:"id"`
	Slug         string   `json:"slug"`
	Diskon       float64  `json:"diskon"`
	Nama         string   `json:"nama"`
	Harga        uint64   `json:"harga"`
	JumlahBeli   uint     `json:"jumlahBeli"`
	Berat        uint     `json:"berat"`
	IsHiasan     bool     `json:"isHiasan"`
	GambarProduk []string `json:"gambarProduk"`
}

func KeranjangProdukFormatter(
	keranjang *[]Keranjang,
) ([]KeranjangProduk, uint64, uint64) {
	var formattedKeranjang []KeranjangProduk
	totalHarga := 0.0
	totalBerat := uint64(0)

	for _, krj := range *keranjang {
		var gambar []string
		for _, gmb := range krj.Produk.GambarProduk {
			gambar = append(gambar, gmb.Nama)
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
			IsHiasan:     krj.Produk.IsHiasan,
			GambarProduk: gambar,
		})
	}
	return formattedKeranjang, uint64(totalHarga), totalBerat
}

func KeranjangFormatter(
	keranjang *[]Keranjang,
	alamatUser []alamat.GetAlamatFormatter,
) KeranjangFormat {
	formattedKeranjang, totalHarga, totalBerat := KeranjangProdukFormatter(keranjang)

	return KeranjangFormat{
		KeranjangProduk: formattedKeranjang,
		TotalHarga:      uint64(totalHarga),
		TotalBerat:      totalBerat,
		Alamat:          alamatUser,
	}

}

func KonfirmasiPesananFormatter(
	user user.ProfilUserFormat,
	keranjang *[]Keranjang,
	alamatUser alamat.GetAlamatFormatter,
	alamatAdmin alamat.GetAlamatFormatter,
	pengiriman []utilities.OngkirFormat,
	metodePembayaran []Metode_Pembayaran,
) KonfirmasiPesananFormat {
	formattedKeranjang, totalHarga, totalBerat := KeranjangProdukFormatter(keranjang)
	return KonfirmasiPesananFormat{
		User:             user,
		AlamatUser:       alamatUser,
		AlamatAdmin:      alamatAdmin,
		KeranjangProduk:  formattedKeranjang,
		Pengiriman:       pengiriman,
		MetodePembayaran: metodePembayaran,
		TotalHarga:       uint64(totalHarga),
		TotalBerat:       totalBerat,
	}
}
