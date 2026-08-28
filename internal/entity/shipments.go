package entity

import (
	"errors"
	"time"
)

type Shipment struct {
	ShipmentID    string
	OrderID       *string
	CarrierID     *string
	Origin        string
	Destination   string
	CurrentStatus string
	Eta           *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

var ErrorInvalid = errors.New("not suitable status")

var validTransitions = map[string][]string{
	"created":          {"picked_up", "canceled"},
	"picked_up":        {"in_transit", "canceled"},
	"in_transit":       {"out_for_delivery", "failed"},
	"out_for_delivery": {"delivered", "failed"},
	"failed":           {"in_transit", "canceled"},
}

func (s *Shipment) CanTransitionTo(newStatus string) bool {
	allowed, ok := validTransitions[s.CurrentStatus]
	if !ok {
		return false
	}

	for _, status := range allowed {
		if status == newStatus {
			return true
		}
	}

	return false
}
