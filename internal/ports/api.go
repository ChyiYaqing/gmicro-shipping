package ports

import (
	"context"

	"github.com/chyiyaqing/gmicro-shipping/internal/application/core/domain"
)

type APIPort interface {
	Create(ctx context.Context, shipping domain.Shipping) (domain.Shipping, error)
}
