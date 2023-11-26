package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	VendorId     uint
	Price        float64
	DeliveryTime uint
	Trip         Trip
	DelayReports []DelayReport
}

func (o Order) IsDelayed() bool {
	return time.Now().After(o.UpdatedAt.Add(time.Duration(o.DeliveryTime) * time.Minute))
}
