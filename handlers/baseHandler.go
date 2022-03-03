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
) Handlers {
	return Handlers{
		authHandler,
		daerahHandler,
		alamatHandler,
		kategoriHandler,
		produkHandler,
	}
}

func (handler Handlers) Setup() {
	for _, h := range handler {
		h.Setup()
	}
}