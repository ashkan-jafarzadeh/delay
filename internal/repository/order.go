package repository

import (
	"context"
	"github.com/ashkan-jafarzadeh/delay/internal/model"
)

type Order interface {
	Create(ctx context.Context, or *model.Order) error
	FindById(ctx context.Context, id uint) (model.Order, error)
	Update(ctx context.Context, order *model.Order) error
}
