package mysql

import (
	"context"
	"github.com/ashkan-jafarzadeh/delay/internal/model"
	"github.com/ashkan-jafarzadeh/delay/internal/repository"
	"gorm.io/gorm"
)

type trip struct {
	db *gorm.DB
}

func NewTrip(db *gorm.DB) repository.Trip {
	return trip{db: db}
}

func (t trip) Create(ctx context.Context, tr *model.Trip) error {
	return t.db.WithContext(ctx).Create(tr).Error
}
