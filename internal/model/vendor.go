package model

import (
	"gorm.io/gorm"
)

type Vendor struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255);not null;"`
	Orders       []Order
	DelayReports []DelayReport
}
