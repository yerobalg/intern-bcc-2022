package alamat

type AlamatInputFormat struct {
	Nama          string `json:"namaAlamat"`
	AlamatLengkap string `json:"alamatLengkap"`
	KodePos       string `json:"kodePos"`
	IDKelurahan   string `json:"idKelurahan"`
	AlamatPembeli bool   `json:"alamatPembeli"`
}

func AlamatInputFormatter(alamat *Alamat) AlamatInputFormat {
	return AlamatInputFormat{
		Nama:          alamat.Nama,
		AlamatLengkap: alamat.AlamatLengkap,
		KodePos:       alamat.KodePos,
		IDKelurahan:   alamat.IDKelurahan,
		AlamatPembeli: alamat.IsUser,
	}
}
