package mysql

import (
	"context"
	"github.com/ashkan-jafarzadeh/delay/internal/model"
	"github.com/ashkan-jafarzadeh/delay/internal/repository"
	"gorm.io/gorm"
)

type agent struct {
	db *gorm.DB
}

func NewAgent(db *gorm.DB) repository.Agent {
	return agent{db: db}
}

func (o agent) Create(ctx context.Context, ag *model.Agent) error {
	return o.db.WithContext(ctx).Create(ag).Error
}

func (o agent) FindWithActiveDelayReportsById(ctx context.Context, id uint) (model.Agent, error) {
	var ag model.Agent
	err := o.db.WithContext(ctx).Preload("DelayReports", func(db *gorm.DB) *gorm.DB {
		return db.Where("reviewed = ?", 0)
	}).Where("id = ?", id).First(&ag).Error

	return ag, err
}
