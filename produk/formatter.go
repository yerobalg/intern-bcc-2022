package produk

type ProdukFormat struct {
	NamaProduk string `json:"nama"`
	Slug       string `json:"slug"`
	Harga      uint64 `json:"harga"`
	Stok       uint   `json:"stok"`
	Deskripsi  string `json:"deskripsi"`
	Gender     string `json:"gender"`
	IsHiasan   bool   `json:"isHiasan"`
	IdKategori uint64 `json:"idKategori"`
	IdTags  	[]uint64 `json:"idTags"`
}

func GetProdukFormat(
	produk *Produk, 
	idKategori uint64, 
	IdTags []uint64,
) ProdukFormat {
	return ProdukFormat{
		NamaProduk: produk.NamaProduk,
		Slug:       produk.Slug,
		Harga:      produk.Harga,
		Stok:       produk.Stok,
		Deskripsi:  produk.Deskripsi,
		Gender:     produk.Gender,
		IsHiasan:   produk.IsHiasan,
		IdKategori: idKategori,
		IdTags: IdTags,
	}
}

// func GetAllProdukFormat(
// 	produk *[]Produk,idKategori []uint64, idTags [][]uint64) []ProdukFormat {
// 	var formattedProduk []ProdukFormat
// 	for _, p := range *produk {
// 		formattedProduk = append(formattedProduk, GetProdukFormat(&p))
// 	}
// 	return formattedProduk
// }
