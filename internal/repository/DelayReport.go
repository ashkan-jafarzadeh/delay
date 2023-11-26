package repository

import (
	"context"
	"github.com/ashkan-jafarzadeh/delay/internal/model"
)

type DelayReport interface {
	Create(ctx context.Context, dr *model.DelayReport) error
	Assign(ctx context.Context, id uint, agentId uint) error
}
