package mysql

import (
	"context"
	"github.com/ashkan-jafarzadeh/delay/internal/model"
	"github.com/ashkan-jafarzadeh/delay/internal/repository"
	"gorm.io/gorm"
)

type order struct {
	db *gorm.DB
}

func NewOrder(db *gorm.DB) repository.Order {
	return order{db: db}
}

func (o order) Create(ctx context.Context, or *model.Order) error {
	return o.db.WithContext(ctx).Create(or).Error
}

func (o order) FindById(ctx context.Context, id uint) (model.Order, error) {
	var or model.Order
	err := o.db.WithContext(ctx).Where("id = ?", id).Preload("Trip").Preload("DelayReports", func(db *gorm.DB) *gorm.DB {
		return db.Where("reviewed = ?", 0)
	}).First(&or).Error

	return or, err
}

func (o order) Update(ctx context.Context, role *model.Order) error {
	return o.db.WithContext(ctx).Where("id = ?", role.ID).Updates(role).Error
}
