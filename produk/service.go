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

func (s *ProdukService) GetBySlug(slug string) (*Produk, error) {
	return s.repo.GetBySlug(slug)
}

func (s *ProdukService) GetByIdSeller(idSeller uint64) (*[]Produk, error) {
	return s.repo.GetByIdSeller(idSeller)
}

func (s *ProdukService) Update(produk *Produk) error {
	return s.repo.Update(produk)
}

func (s *ProdukService) Delete(produk *Produk) error {
	return s.repo.Delete(produk)
}
