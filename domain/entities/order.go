package entities

import (
	"time"
)

type Order struct {
	Id            int         `gorm:"primaryKey" json:"id"`              // Set primary key with auto-increment
	O_status      string      `gorm:"not null" json:"o_status"`          // Not-null constraint
	O_timestamp   time.Time   `gorm:"not null" json:"o_timestamp"`       // Not-null constraint
	O_total_price float64     `gorm:"not null" json:"o_total_price"`     // Not-null constraint
	UserID        int         `gorm:"not null" json:"userID"`            // Foreign key with not-null
	User          User        `gorm:"foreignKey:UserID"`                 // Foreign key association
	ShipmentID    *int        `gorm:"default:null" json:"shipmentID"`    // Nullable foreign key
	Shipment      Shipment    `gorm:"foreignKey:ShipmentID"`             // Foreign key association
	TransactionID *int        `gorm:"default:null" json:"transactionID"` // Nullable foreign key
	Transaction   Transaction `gorm:"foreignKey:TransactionID"`          // Foreign key association
}
