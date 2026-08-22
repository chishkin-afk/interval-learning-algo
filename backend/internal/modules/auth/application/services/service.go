package services

import (
	"context"
	"errors"
	"log/slog"
	"uuid"

	"github.com/chishkin-afk/interval-learning-algo/backend/internal/application/dtos/requests"
	"github.com/chishkin-afk/interval-learning-algo/backend/internal/application/dtos/responses"
	"github.com/chishkin-afk/interval-learning-algo/backend/internal/common/config"
	"github.com/chishkin-afk/interval-learning-algo/backend/internal/modules/auth/domain/user"
)

var (
	ErrInternalServer = errors.New("internal server error")
)

type AuthService struct {
	cfg             *config.Config
	log             *slog.Logger
	userPersistence user.UserPersistenceRepository
}

func New(
	cfg *config.Config,
	log *slog.Logger,
	userPersistence user.UserPersistenceRepository,
) *AuthService {
	return &AuthService{
		cfg:             cfg,
		log:             log,
		userPersistence: userPersistence,
	}
}

func (as *AuthService) Register(ctx context.Context, req *requests.Register) (*responses.Token, error) {
	panic("not impl")
}

func (as *AuthService) Login(ctx context.Context, req *requests.Login) (*responses.Token, error) {
	panic("not impl")
}

func (as *AuthService) GetByID(ctx context.Context, id uuid.UUID) (*responses.User, error) {
	panic("not impl")
}

func (as *AuthService) GetSelf(ctx context.Context) (*responses.User, error) {
	panic("not impl")
}

func (as *AuthService) Update(ctx context.Context, req *requests.UpdateUser) (*responses.User, error) {
	panic("not impl")
}

func (as *AuthService) Delete(ctx context.Context) error {
	panic("not impl")
}

func handleError(err error) *ServiceError {
	switch {
	case errors.Is(err, context.DeadlineExceeded),
		errors.Is(err, context.Canceled):
		return NewServiceError(KindTimeout, err)
	case errors.Is(err, user.ErrAlreadyExists):
		return NewServiceError(KindConflict, err)
	case errors.Is(err, user.ErrNotFound):
		return NewServiceError(KindNotFound, err)
	}

	return NewServiceError(KindInternal, ErrInternalServer)
}
