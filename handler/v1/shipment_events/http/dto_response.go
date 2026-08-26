package shipmentevent

import (
	"time"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
)

type ShipmentEventObject struct {
	ShipmentEventID string
	ShipmentID      string
	Status          string
	Location        *string
	Notes           *string
	OccuredAt       time.Time
	CreatedAt       time.Time
}

func (ShipmentEventObject) ParseFromEntity(e entity.ShipmentEvent) ShipmentEventObject {
	return ShipmentEventObject{
		ShipmentEventID: e.ShipmentEventID,
		ShipmentID:      e.ShipmentID,
		Status:          e.Status,
		Location:        e.Location,
		Notes:           e.Notes,
		OccuredAt:       e.OccurredAt,
		CreatedAt:       e.CreatedAt,
	}
}
