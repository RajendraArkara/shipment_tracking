package repository

import (
	"context"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
)

type ShipmentEventRepository interface {
	GetByShipmentID(ctx context.Context, ShipmentEventID string) ([]entity.ShipmentEvent, error)
}
