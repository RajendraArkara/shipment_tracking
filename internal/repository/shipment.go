package repository

import (
	"context"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
)

type ShipmentRepository interface {
	GetAll(ctx context.Context) ([]entity.Shipment, error)
}
