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
	return s.repo.GetBySlug(slug)
}

// func (s *ProdukService) GetByIdSeller(idSeller uint64) (*[]Produk, error) {
// 	return s.repo.GetByIdSeller(idSeller)
// }

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
	return s.repo.DeleteGambarProduk(idProduk, nama);
}

func (s *ProdukService) Delete(produk *Produk) error {
	return s.repo.Delete(produk)
}
