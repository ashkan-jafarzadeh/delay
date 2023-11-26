package repository

import (
	"context"
	"github.com/ashkan-jafarzadeh/delay/internal/aggregate"
	"github.com/ashkan-jafarzadeh/delay/internal/model"
	"time"
)

type Vendor interface {
	Create(ctx context.Context, tr *model.Vendor) error
	Report(ctx context.Context, from time.Time) ([]aggregate.VendorReport, error)
}
