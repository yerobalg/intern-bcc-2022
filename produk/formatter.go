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
}

func GetProdukFormat(produk *Produk) ProdukFormat {
	return ProdukFormat{
		NamaProduk: produk.NamaProduk,
		Slug:       produk.Slug,
		Harga:      produk.Harga,
		Stok:       produk.Stok,
		Deskripsi:  produk.Deskripsi,
		Gender:     produk.Gender,
		IsHiasan:   produk.IsHiasan,
		// IdKategori: produk.IDKategori,
	}
}

func GetAllProdukFormat(produk *[]Produk) []ProdukFormat {
	var formattedProduk []ProdukFormat
	for _, p := range *produk {
		formattedProduk = append(formattedProduk, GetProdukFormat(&p))
	}
	return formattedProduk
}
