package author

import (
	"ca-library-app/internal/adapters/api"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	authorURL  = "/authors/:author_id"
	authorsURL = "/authors"
)

type handler struct {
	authorService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{authorService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(authorsURL, h.GetAll)
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)
}
