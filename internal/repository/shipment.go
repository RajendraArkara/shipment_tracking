package repository

import (
	"context"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
)

type ShipmentRepository interface {
	GetAll(ctx context.Context) ([]entity.Shipment, error)
	GetByID(ctx context.Context, id string) (*entity.Shipment, error)
	Create(ctx context.Context, data *entity.Shipment) (*entity.Shipment, error)
}
