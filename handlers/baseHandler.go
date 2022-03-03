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
) Handlers {
	return Handlers{
		authHandler,
		daerahHandler,
		alamatHandler,
		kategoriHandler,
	}
}

func (handler Handlers) Setup() {
	for _, h := range handler {
		h.Setup()
	}
}