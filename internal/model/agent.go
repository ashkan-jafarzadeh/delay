package model

import (
	"gorm.io/gorm"
)

type Agent struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255);not null;"`
	DelayReports []DelayReport
}
