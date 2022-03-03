package alamat

type AlamatInputFormat struct {
	Nama          string `json:"namaAlamat"`
	AlamatLengkap string `json:"alamatLengkap"`
	KodePos       string `json:"kodePos"`
	IDKelurahan   string `json:"idKelurahan"`
	AlamatPembeli bool   `json:"alamatPembeli"`
}

type GetAlamatFormatter struct {
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

func GetUserAlamatFormatter(alamat []Alamat) []GetAlamatFormatter {
	var alamatFormatted []GetAlamatFormatter
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
		alamatFormatted = append(alamatFormatted, GetAlamatFormatter{
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

func GetAlamatByIdFormatter(alamat Alamat) GetAlamatFormatter {
	kelurahan := DaerahFormat{
		ID:   alamat.Kelurahan.IDKel,
		Nama: alamat.Kelurahan.Nama,
	}
	kecamatan := DaerahFormat{
		ID:   alamat.Kelurahan.Kecamatan.IDKec,
		Nama: alamat.Kelurahan.Kecamatan.Nama,
	}
	kabupatenKota := DaerahFormat{
		ID:   alamat.Kelurahan.Kecamatan.Kabupaten.IDKab,
		Nama: alamat.Kelurahan.Kecamatan.Kabupaten.Nama,
	}
	provinsi := DaerahFormat{
		ID:   alamat.Kelurahan.Kecamatan.Kabupaten.Provinsi.IDProv,
		Nama: alamat.Kelurahan.Kecamatan.Kabupaten.Provinsi.Nama,
	}
	return GetAlamatFormatter{
		ID:            alamat.ID,
		Nama:          alamat.Nama,
		AlamatLengkap: alamat.AlamatLengkap,
		KodePos:       alamat.KodePos,
		Kelurahan:     kelurahan,
		Kecamatan:     kecamatan,
		KabupatenKota: kabupatenKota,
		Provinsi:      provinsi,
	}
}
