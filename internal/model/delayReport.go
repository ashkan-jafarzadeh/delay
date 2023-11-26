package model

import (
	"gorm.io/gorm"
)

type DelayReport struct {
	gorm.Model
	AgentId  uint `gorm:"default:NULL"`
	OrderId  uint
	VendorId uint
	Reviewed bool
}
