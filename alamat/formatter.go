package alamat

type AlamatInputFormat struct {
	Nama          string `json:"namaAlamat"`
	AlamatLengkap string `json:"alamatLengkap"`
	KodePos       string `json:"kodePos"`
	IDKelurahan   string `json:"idKelurahan"`
	AlamatPembeli bool   `json:"alamatPembeli"`
}

type GetUserAlamatFormat struct {
	ID            uint         `json:"id"`
	Nama          string       `json:"nama"`
	AlamatLengkap string       `json:"alamatLengkap"`
	KodePos       string       `json:"kodePos"`
	Kelurahan     DaerahFormat `json:"kelurahan"`
	Kecamatan     DaerahFormat `json:"kecamatan"`
	KabupatenKota DaerahFormat `json:"kabupatenKota"`
	Provinsi      DaerahFormat `json:"provinsi"`
}

type DaerahFormat struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
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

func GetUserAlamatFormatter(alamat []Alamat) []GetUserAlamatFormat {
	var alamatFormatted []GetUserAlamatFormat
	for _, almt := range alamat {
		kelurahan := DaerahFormat{
			ID:   almt.Kelurahan.IDKel,
			Nama: almt.Kelurahan.Nama,
		}
		kecamatan := DaerahFormat{
			ID:   almt.Kelurahan.Kecamatan.IDKec,
			Nama: almt.Kelurahan.Kecamatan.Nama,
		}
		kabupatenKota := DaerahFormat{
			ID:   almt.Kelurahan.Kecamatan.Kabupaten.IDKab,
			Nama: almt.Kelurahan.Kecamatan.Kabupaten.Nama,
		}
		provinsi := DaerahFormat{
			ID:   almt.Kelurahan.Kecamatan.Kabupaten.Provinsi.IDProv,
			Nama: almt.Kelurahan.Kecamatan.Kabupaten.Provinsi.Nama,
		}
		alamatFormatted = append(alamatFormatted, GetUserAlamatFormat{
			ID:            almt.ID,
			Nama:          almt.Nama,
			AlamatLengkap: almt.AlamatLengkap,
			KodePos:       almt.KodePos,
			Kelurahan:     kelurahan,
			Kecamatan:     kecamatan,
			KabupatenKota: kabupatenKota,
			Provinsi:      provinsi,
		})
	}
	return alamatFormatted
}
