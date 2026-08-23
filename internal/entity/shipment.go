package entity

import "time"

type Shipment struct {
	ShipmentID    string
	OrderID       string
	CarrierID     string
	Origin        string
	Destination   string
	CurrentStatus string
	Eta           *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
