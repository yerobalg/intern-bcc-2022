package produk

type ProdukService struct {
	repo ProdukRepository
}

func NewProdukService(produkRepo ProdukRepository) ProdukService {
	return ProdukService{repo: produkRepo}
}
