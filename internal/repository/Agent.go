package repository

import (
	"context"
	"github.com/ashkan-jafarzadeh/delay/internal/model"
)

type Agent interface {
	Create(ctx context.Context, ag *model.Agent) error
	FindWithActiveDelayReportsById(ctx context.Context, id uint) (model.Agent, error)
}
