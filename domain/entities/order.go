package entities

import (
	"time"
)

type Order struct {
	Id            int       `json:"id"`
	O_status      string    `json:"o_status"`
	O_timestamp   time.Time `json:"o_timestamp"`
	O_total_price float64   `json:"o_total_price"`
	UserID        int       `json:"userID"`
	// User          User        `gorm:"foreignKey:UserID"`
	ShipmentID int `json:"shipmentID"`
	// Shipment      Shipment    `gorm:"foreignKey:ShipmentID"`
	TransactionID int `json:"transactionID"`
	// Transaction   Transaction `gorm:"foreignKey:TransactionID"`
}
