package entity

import "time"

type Webhook struct {
	ID         string
	ShipmentID string
	TargetUrl  string
	SecretKey  string
	IsActive   bool
	CreatedAt  time.Time
}
