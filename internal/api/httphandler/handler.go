package httphandler

import (
	"fmt"
	"net/http"

	api "github.com/aliocode/relic-example/api/http"
)

var (
	_ api.ServerInterface = (*Handler)(nil)
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (s *Handler) GetUserId(w http.ResponseWriter, r *http.Request, id int) {
	_ = r
	_, _ = w.Write([]byte(fmt.Sprintf("received GetUserId: %d", id)))
	w.WriteHeader(http.StatusOK)
}
