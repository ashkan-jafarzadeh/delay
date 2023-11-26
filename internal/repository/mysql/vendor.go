package mysql

import (
	"context"
	"github.com/ashkan-jafarzadeh/delay/internal/aggregate"
	"github.com/ashkan-jafarzadeh/delay/internal/model"
	"github.com/ashkan-jafarzadeh/delay/internal/repository"
	"gorm.io/gorm"
	"time"
)

type vendor struct {
	db *gorm.DB
}

func NewVendor(db *gorm.DB) repository.Vendor {
	return vendor{db: db}
}

func (v vendor) Create(ctx context.Context, tr *model.Vendor) error {
	return v.db.WithContext(ctx).Create(tr).Error
}

func (v vendor) Report(ctx context.Context, from time.Time) ([]aggregate.VendorReport, error) {
	var data []aggregate.VendorReport
	err := v.db.WithContext(ctx).Raw(`SELECT vendors.id, vendors.name, count(delay_reports.id) as delay_count
FROM vendors
JOIN delay_reports on vendors.id = delay_reports.vendor_id
WHERE delay_reports.created_at >= ?
GROUP BY vendors.id
order by delay_count DESC`, from).Scan(&data).Error

	return data, err
}
