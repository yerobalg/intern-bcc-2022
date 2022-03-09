package pesanan

import (
	"clean-arch-2/alamat"
	"clean-arch-2/keranjang"
	"gorm.io/gorm"
	"time"
)

type InputPesanan struct {
	MetodePembayaran string   `json:"metodePembayaran" binding:"required"`
	KodeKurir        string   `json:"kodeKurir" binding:"required"`
	KodeLayanan      string   `json:"kodeLayanan" binding:"required"`
	HargaOngkosKirim uint64   `json:"hargaOngkosKirim"`
	EstimasiKirim    string   `json:"estimasiKirim" binding:"required"`
	IDAlamat         uint64   `json:"idAlamat"`
	IDKeranjang      []uint64 `json:"idKeranjang"`
}

type Pesanan struct {
	gorm.Model
	MetodePembayaran string              `json:"metodePembayaran" gorm:"type:varchar(20);not null"`
	KodeKurir        string              `json:"kodeKurir" gorm:"type:varchar(20);not null"`
	KodeLayanan      string              `json:"kodeLayanan" gorm:"type:varchar(20);not null"`
	HargaOngkosKirim uint64              `json:"hargaOngkosKirim" gorm:"type:bigint;not null"`
	EstimasiKirim    string              `json:"estimasiKirim" gorm:"type:varchar(20);not null"`
	BiayaAdmin       uint64              `json:"biayaAdmin" gorm:"type:bigint;not null"`
	StatusPemesanan  string              `json:"statusPemesanan" gorm:"type:varchar(50);default:'belum dibayar'"`
	TanggalBayar     time.Time           `json:"tanggalBayar" gorm:"type:timestamp;"`
	IDAlamat         uint64              `json:"idAlamat" gorm:"type:bigint;not null"`
	Alamat           alamat.Alamat       `json:"alamat" gorm:"foreignkey:IDAlamat;constraint:OnDelete:CASCADE"`
	Keranjang        []Keranjang_Pesanan `json:"keranjang" gorm:"foreignkey:IDPesanan;constraint:OnDelete:CASCADE"`
}

type Keranjang_Pesanan struct {
	IDPesanan   uint64              `json:"idPesanan" gorm:"type:bigint;not null"`
	IDKeranjang keranjang.Keranjang `json:"idKeranjang" gorm:"type:bigint;not null"`
	Pesanan     Pesanan             `json:"pesanan" gorm:"foreignkey:IDPesanan;constraint:OnDelete:CASCADE"`
	Keranjang   keranjang.Keranjang `json:"keranjang" gorm:"foreignkey:IDKeranjang;constraint:OnDelete:CASCADE"`
}
