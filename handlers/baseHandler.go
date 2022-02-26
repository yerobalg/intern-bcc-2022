package handlers

type Handlers []Handler

type Handler interface {
	Setup()
}

func NewHandlers(
	authHandler AuthHandler,
	daerahHandler DaerahHandler,
) Handlers {
	return Handlers{
		authHandler,
		daerahHandler,

	}
}

func (handler Handlers) Setup() {
	for _, h := range handler {
		h.Setup()
	}
}