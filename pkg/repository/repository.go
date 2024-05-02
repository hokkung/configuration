package repository

import (
	"context"
)

type Entity[K any] interface {
	EntID() K
}

type Repository[E Entity[K], K any] interface {
	Create(ctx context.Context, ent *E) error
	Update(ctx context.Context, ent *E) error
	FindByID(ctx context.Context, id interface{}) (*E, error)
	Delete(ctx context.Context, ent E) error
}
