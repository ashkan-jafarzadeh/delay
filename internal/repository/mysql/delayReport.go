package mysql

import (
	"context"
	"github.com/ashkan-jafarzadeh/delay/internal/model"
	"gorm.io/gorm"
)

type delayReport struct {
	db *gorm.DB
}

func NewDelayReport(db *gorm.DB) *delayReport {
	return &delayReport{db: db}
}

func (d delayReport) Create(ctx context.Context, dr *model.DelayReport) error {
	return d.db.WithContext(ctx).Create(dr).Error
}

func (d delayReport) Assign(ctx context.Context, id uint, agentId uint) error {
	var dr model.DelayReport
	return d.db.WithContext(ctx).Model(&dr).Where("id = ?", id).UpdateColumn("agent_id", agentId).Error
}
