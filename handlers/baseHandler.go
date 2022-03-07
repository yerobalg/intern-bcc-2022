package handlers

type Handlers []Handler

type Handler interface {
	Setup()
}

func NewHandlers(
	authHandler AuthHandler,
	daerahHandler DaerahHandler,
	alamatHandler AlamatHandler,
	kategoriHandler KategoriHandler,
	produkHandler ProdukHandler,
	keranjangHandler KeranjangHandler,
) Handlers {
	return Handlers{
		authHandler,
		daerahHandler,
		alamatHandler,
		kategoriHandler,
		produkHandler,
		keranjangHandler,
	}
}

func (handler Handlers) Setup() {
	for _, h := range handler {
		h.Setup()
	}
}