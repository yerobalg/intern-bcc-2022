package alamat

type AlamatInputFormat struct {
	Nama          string `json:"namaAlamat"`
	AlamatLengkap string `json:"alamatLengkap"`
	KodePos       string `json:"kodePos"`
	IDKabupaten   string `json:"idKabupaten"`
	AlamatPembeli bool   `json:"alamatPembeli"`
}

type KonfirmasiAlamatFormat struct {

}

type GetAlamatFormatter struct {
	ID            uint            `json:"id"`
	Nama          string          `json:"nama"`
	AlamatLengkap string          `json:"alamatLengkap"`
	KodePos       string          `json:"kodePos"`
	KabupatenKota KabupatenFormat `json:"kabupatenKota"`
	Provinsi      ProvinsiFormat  `json:"provinsi"`
}

type ProvinsiFormat struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
}

type KabupatenFormat struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
	Tipe string `json:"tipe"`
}

func AlamatInputFormatter(alamat *Alamat) AlamatInputFormat {
	return AlamatInputFormat{
		Nama:          alamat.Nama,
		AlamatLengkap: alamat.AlamatLengkap,
		KodePos:       alamat.KodePos,
		IDKabupaten:   alamat.IDKabupaten,
		AlamatPembeli: alamat.IsUser,
	}
}

func GetUserAlamatFormatter(alamat *[]Alamat) []GetAlamatFormatter {
	var alamatFormatted []GetAlamatFormatter
	for _, almt := range *alamat {
		alamatFormatted = append(alamatFormatted, GetAlamatByIdFormatter(almt))
	}
	return alamatFormatted
}

func GetAlamatByIdFormatter(alamat Alamat) GetAlamatFormatter {
	kabupatenKota := KabupatenFormat{
		ID:   alamat.Kabupaten.IDKab,
		Nama: alamat.Kabupaten.Nama,
		Tipe: alamat.Kabupaten.Tipe,
	}
	provinsi := ProvinsiFormat{
		ID:   alamat.Kabupaten.Provinsi.IDProv,
		Nama: alamat.Kabupaten.Provinsi.Nama,
	}
	return GetAlamatFormatter{
		ID:            alamat.ID,
		Nama:          alamat.Nama,
		AlamatLengkap: alamat.AlamatLengkap,
		KodePos:       alamat.KodePos,
		KabupatenKota: kabupatenKota,
		Provinsi:      provinsi,
	}
}
