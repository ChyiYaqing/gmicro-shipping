package ports

import (
	"context"

	"github.com/chyiyaqing/gmicro-shipping/internal/application/core/domain"
)

type DBPort interface {
	Get(ctx context.Context, id string) (domain.Shipping, error)
	Save(ctx context.Context, shipping *domain.Shipping) error
}
