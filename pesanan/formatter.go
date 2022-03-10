package pesanan

import (
	"clean-arch-2/alamat"
	"clean-arch-2/keranjang"
	"clean-arch-2/user"
	"time"
)

type PesananInputFormat struct {
	ID                 uint64   `json:"id"`
	KodeKurir          string   `json:"kodeKurir"`
	OngkosKirim        uint64   `json:"hargaOngkosKirim"`
	EstimasiKirim      string   `json:"estimasiKirim"`
	IDMetodePembayaran uint64   `json:"idMetodePembayaran"`
	IDAlamat           uint64   `json:"idAlamat"`
	IDKeranjang        []uint64 `json:"idKeranjang"`
}

type PesananOutputFormat struct {
	ID               uint64                      `json:"id"`
	User             user.ProfilUserFormat       `json:"user"`
	AlamatUser       alamat.GetAlamatFormatter   `json:"alamatUser"`
	AlamatAdmin      alamat.GetAlamatFormatter   `json:"alamatSeller"`
	KodeKurir        string                      `json:"kodeKurir"`
	OngkosKirim      uint64                      `json:"hargaOngkosKirim"`
	EstimasiKirim    string                      `json:"estimasiKirim"`
	IDMetode         uint64                      `json:"idMetode"`
	MetodePembayaran string                      `json:"metodePembayaran"`
	BiayaAdmin       uint64                      `json:"biayaAdmin"`
	BatasBayar       time.Time                   `json:"batasBayar"`
	TanggalBayar     time.Time                   `json:"tanggalBayar"`
	StatusPemesanan  string                      `json:"statusPemesanan"`
	Produk           []keranjang.KeranjangProduk `json:"produk"`
	TotalBeratProduk uint64                      `json:"totalBeratProduk"`
	TotalHarga       uint64                      `json:"totalHarga"`
}

func PesananInputFormatter(
	pesanan *Pesanan,
	IDKeranjang []uint64,
) PesananInputFormat {
	return PesananInputFormat{
		ID:                 uint64(pesanan.ID),
		KodeKurir:          pesanan.KodeKurir,
		OngkosKirim:        pesanan.HargaOngkosKirim,
		EstimasiKirim:      pesanan.EstimasiKirim,
		IDMetodePembayaran: uint64(pesanan.IDMetode),
		IDAlamat:           uint64(pesanan.IDAlamat),
		IDKeranjang:        IDKeranjang,
	}
}

func PesananOutputFormatter(
	pesananObj *Pesanan,
	alamatAdmin alamat.Alamat,
) PesananOutputFormat {

	keranjangObj := pesananObj.Keranjang_Pesanan
	var keranjangProduk []keranjang.Keranjang
	for _, k := range keranjangObj {
		keranjangProduk = append(keranjangProduk, k.Keranjang)
	}

	formattedProduk, totalHargaProduk, totalBeratProduk := keranjang.KeranjangProdukFormatter(&keranjangProduk)

	return PesananOutputFormat{
		ID:               uint64(pesananObj.ID),
		User:             user.ProfilUserFormatter(&pesananObj.User),
		AlamatUser:       alamat.GetAlamatByIdFormatter(pesananObj.Alamat),
		AlamatAdmin:      alamat.GetAlamatByIdFormatter(alamatAdmin),
		KodeKurir:        pesananObj.KodeKurir,
		OngkosKirim:      pesananObj.HargaOngkosKirim,
		EstimasiKirim:    pesananObj.EstimasiKirim,
		IDMetode:         pesananObj.IDMetode,
		MetodePembayaran: pesananObj.Metode.Nama,
		BiayaAdmin:       pesananObj.Metode.BiayaAdmin,
		BatasBayar:       pesananObj.BatasBayar,
		TanggalBayar:     pesananObj.TanggalBayar,
		StatusPemesanan:  pesananObj.StatusPemesanan,
		Produk:           formattedProduk,
		TotalBeratProduk: totalBeratProduk,
		TotalHarga:       totalHargaProduk + pesananObj.HargaOngkosKirim + pesananObj.Metode.BiayaAdmin,
	}
}
