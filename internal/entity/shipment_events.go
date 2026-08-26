package entity

import "time"

type ShipmentEvent struct {
	ShipmentEventID string
	ShipmentID      string
	Status          string
	Location        *string
	Notes           *string
	OccurredAt      time.Time
	CreatedAt       time.Time
}
