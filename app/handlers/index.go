package handlers

type Handlers []Handler

type Handler interface {
	Setup()
}

func NewHandlers(
	postHandler PostHandler,
) Handlers {
	return Handlers{
		postHandler,
	}
}

func (handler Handlers) Setup() {
	for _, h := range handler {
		h.Setup()
	}
}
