package produk

type ProdukService struct {
	repo ProdukRepository
}

func NewProdukService(produkRepo ProdukRepository) ProdukService {
	return ProdukService{repo: produkRepo}
}

func (s *ProdukService) Save(produk *Produk) error {
	return s.repo.Save(produk)
}

func (s *ProdukService) GetById(id uint64) (*Produk, error) {
	return s.repo.GetById(id)
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
