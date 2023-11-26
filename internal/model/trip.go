package model

import (
	"gorm.io/gorm"
)

type Trip struct {
	gorm.Model
	OrderId uint   `gorm:"foreignkey:OrderId"`
	Status  Status `gorm:"type:enum('ASSIGNED', 'AT_VENDOR', 'PICKED', 'DELIVERED');default:'ASSIGNED'"`
}

// Status type with constants
type Status string

const (
	TripAssigned  Status = "ASSIGNED"
	TripAtVendor  Status = "AT_VENDOR"
	TripPicked    Status = "PICKED"
	TripDelivered Status = "DELIVERED"
)

func (t Trip) IsActive() bool {
	return t.Status == TripAssigned || t.Status == TripAtVendor || t.Status == TripPicked
}

func TripStatuses() []Status {
	return []Status{
		TripAssigned,
		TripAtVendor,
		TripPicked,
		TripDelivered,
	}
}
