package httphandler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	api "github.com/aliocode/relic-example/api/http"
	"github.com/aliocode/relic-example/internal/domain/user"
)

var (
	_ api.ServerInterface = (*Handler)(nil)
)

type UserService interface {
	CreateUser(ctx context.Context, name string, email string) (user.User, error)
	FetchByEmail(ctx context.Context, email string) (user.User, error)
}

type Handler struct {
	userService UserService
}

func New(userService UserService) *Handler {
	return &Handler{userService: userService}
}

func (s *Handler) GetUsersEmail(w http.ResponseWriter, r *http.Request, email string) {
	_ = r

	res, err := s.userService.FetchByEmail(context.Background(), email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := struct {
		ID        string
		Name      string
		Email     string
		CreatedAt time.Time
		UpdatedAT time.Time
	}{res.ID, res.Name, res.Email, res.CreatedAt, res.UpdatedAt}

	bytes, err := json.Marshal(&response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, _ = w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func (s *Handler) PostUsers(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.WriteHeader(http.StatusOK)
}
