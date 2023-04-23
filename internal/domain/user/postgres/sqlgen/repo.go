package sqlgen

import (
	"context"

	"github.com/aliocode/relic-example/internal/domain/user"
	"github.com/aliocode/relic-example/internal/domain/user/service"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ service.Repo = (*Repo)(nil)

type Repo struct {
	db *Queries
}

func NewRepo(pool *pgxpool.Pool) *Repo {
	return &Repo{db: New(pool)}
}

func (repo *Repo) CreateUser(ctx context.Context, in user.User) error {
	err := repo.db.CreateUser(ctx, CreateUserParams{
		ID:    in.ID,
		Name:  in.Name,
		Email: in.Email,
	})
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repo) FetchByEmail(ctx context.Context, email string) (user.User, error) {
	res, err := repo.db.FetchByEmail(ctx, email)
	if err != nil {
		return user.User{}, err
	}
	return user.User{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}
