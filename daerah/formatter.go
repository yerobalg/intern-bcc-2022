package daerah

type OutputProvinsi struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
}

type OutputKabupaten struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
	Tipe string `json:"tipe"`
}

func ProvinsiFormat(prov *[]Provinsi) []OutputProvinsi {
	var provinsi []OutputProvinsi
	for _, prov := range *prov {
		provinsi = append(provinsi, OutputProvinsi{
			ID:   prov.IDProv,
			Nama: prov.Nama,
		})
	}
	return provinsi
}

func KabupatenFormat(kab *[]Kabupaten) []OutputKabupaten {
	var kabupaten []OutputKabupaten
	for _, kab := range *kab {
		kabupaten = append(kabupaten, OutputKabupaten{
			ID:   kab.IDKab,
			Nama: kab.Nama,
			Tipe: kab.Tipe,
		})
	}
	return kabupaten
}
