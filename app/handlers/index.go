package handlers

type Handlers []Handler

type Handler interface {
	Setup()
}

func NewHandlers(
	authHandler AuthHandler,
) Handlers {
	return Handlers{
		authHandler,
	}
}

func (handler Handlers) Setup() {
	for _, h := range handler {
		h.Setup()
	}
}
