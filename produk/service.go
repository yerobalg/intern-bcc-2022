package produk

type ProdukService struct {
	repo ProdukRepository
}

func NewProdukService(produkRepo ProdukRepository) ProdukService {
	return ProdukService{repo: produkRepo}
}

func (s *ProdukService) Save(produk *Produk, idKategori []uint64) error {
	err := s.repo.Save(produk)
	if err != nil {
		return err
	}
	err = s.repo.SaveKategoriProduk(idKategori, uint64(produk.ID))
	return err
}

func (s *ProdukService) SaveGambar(
	idProduk uint64,
	nama []string,
) error {
	return s.repo.SaveGambarProduk(idProduk, nama)
}

func (s *ProdukService) GetBySlug(slug string) (*Produk, error) {
	res, err := s.repo.GetBySlug(slug)
	return res, err
}

func (s *ProdukService) GetHiasanProduk(idKategori uint64) ([]ProdukLite, error, int64) {
	return s.repo.GetHiasanProduk(idKategori)
}

func (s *ProdukService) Update(
	produk *Produk,
	idKategori []uint64,
) error {
	if err := s.repo.Update(produk); err != nil {
		return err
	}

	if err := s.repo.DeleteKategoriProduk(uint64(produk.ID)); err != nil {
		return err
	}

	return s.repo.SaveKategoriProduk(idKategori, uint64(produk.ID))
}

func (s *ProdukService) DeleteGambarProduk(
	idProduk uint64,
	nama string,
) error {
	return s.repo.DeleteGambarProduk(idProduk, nama)
}

func (s *ProdukService) CariProduk(
	idKategori []uint64,
	kataKunci string,
	maksHarga uint64,
	gender string,
	page int,
	sort string,
) ([]ProdukLite, error, int64) {
	if sort == "popular" {
		return s.repo.CariProdukTerlaris(
			idKategori,
			kataKunci,
			maksHarga,
			gender,
			page,
		)
	} else {
		return s.repo.CariProdukTerbaru(
			idKategori,
			kataKunci,
			maksHarga,
			gender,
			page,
		)
	}
}

func (s *ProdukService) Delete(produk *Produk) error {
	return s.repo.Delete(produk)
}

func (s *ProdukService) GetGambarProduk(idProduk []uint64) ([]Produk, error) {
	return s.repo.GetGambarProduk(idProduk)
}

func (s *ProdukService) KurangiStok(
	idProduk []uint64,
	jumlah []uint,
) error {
	return s.repo.KurangiStok(idProduk, jumlah)
}
