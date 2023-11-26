package repository

import (
	"context"
	"github.com/ashkan-jafarzadeh/delay/internal/model"
)

type Trip interface {
	Create(ctx context.Context, tr *model.Trip) error
}
