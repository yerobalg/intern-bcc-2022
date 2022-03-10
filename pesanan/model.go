package pesanan

import (
	"clean-arch-2/alamat"
	"clean-arch-2/keranjang"
	"clean-arch-2/user"
	"gorm.io/gorm"
	"time"
)

type InputPesanan struct {
	KodeKurir          string   `json:"kodeKurir" binding:"required"`
	KodeLayanan        string   `json:"kodeLayanan" binding:"required"`
	HargaOngkosKirim   uint64   `json:"hargaOngkosKirim"`
	EstimasiKirim      string   `json:"estimasiKirim" binding:"required"`
	IDMetodePembayaran uint64   `json:"idMetodePembayaran"`
	IDAlamat           uint64   `json:"idAlamat"`
	IDKeranjang        []uint64 `json:"idKeranjang"`
}

type Pesanan struct {
	gorm.Model
	KodeKurir         string                      `json:"kodeKurir" gorm:"type:varchar(20);not null"`
	KodeLayanan       string                      `json:"kodeLayanan" gorm:"type:varchar(20);not null"`
	HargaOngkosKirim  uint64                      `json:"hargaOngkosKirim" gorm:"type:bigint;not null"`
	EstimasiKirim     string                      `json:"estimasiKirim" gorm:"type:varchar(20);not null"`
	StatusPemesanan   string                      `json:"statusPemesanan" gorm:"type:varchar(50);default:'belum dibayar'"`
	BatasBayar        time.Time                   `json:"BatasBayar" gorm:"type:timestamp;"`
	TanggalBayar      time.Time                   `json:"tanggalBayar" gorm:"type:timestamp;"`
	IDUser            uint64                      `json:"idUser" gorm:"type:bigint;not null"`
	IDMetode          uint64                      `json:"idMetode" gorm:"type:bigint;not null"`
	IDAlamat          uint64                      `json:"idAlamat" gorm:"type:bigint;not null"`
	User              user.Users                  `json:"user" gorm:"foreignkey:IDUser;constraint:OnDelete:CASCADE"`
	Alamat            alamat.Alamat               `json:"alamat" gorm:"foreignkey:IDAlamat;constraint:OnDelete:CASCADE"`
	Metode            keranjang.Metode_Pembayaran `json:"metode" gorm:"foreignkey:IDMetode;constraint:OnDelete:CASCADE"`
	Keranjang_Pesanan []Keranjang_Pesanan         `json:"keranjang" gorm:"foreignkey:IDPesanan;constraint:OnDelete:CASCADE"`
}

func (Pesanan) TableName() string {
	return "pesanan"
}

type Keranjang_Pesanan struct {
	IDPesanan   uint64              `json:"idPesanan" gorm:"type:bigint;not null"`
	IDKeranjang uint64              `json:"idKeranjang" gorm:"type:bigint;not null"`
	Pesanan     Pesanan             `json:"pesanan" gorm:"foreignkey:IDPesanan;constraint:OnDelete:CASCADE"`
	Keranjang   keranjang.Keranjang `json:"keranjang" gorm:"foreignkey:IDKeranjang;constraint:OnDelete:CASCADE"`
}

func (Keranjang_Pesanan) TableName() string {
	return "keranjang_pesanan"
}
