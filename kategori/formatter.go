package kategori

type KategoriFormat struct {
	ID   uint   `json:"id"`
	Nama string   `json:"nama"`
}

func GetKategoriFormatter(kategori *Kategori) KategoriFormat {
	return KategoriFormat{
		ID:   kategori.ID,
		Nama: kategori.Nama,
	}
}

func GetSemuaKategoriFormatter(kategori *[]Kategori) []KategoriFormat {
	formattedKategori := []KategoriFormat{}
	for _, ktg := range *kategori {
		formattedKategori = append(formattedKategori, GetKategoriFormatter(&ktg))
	}
	return formattedKategori
}