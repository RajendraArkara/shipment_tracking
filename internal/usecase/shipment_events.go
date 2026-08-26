package usecase

import (
	"context"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
	"github.com/RajendraArkara/shipment_tracking/internal/repository"
)

type IShipmentEventsUseCase interface {
	GetByShipmentID(ctx context.Context, ShipmentEventID string) ([]entity.ShipmentEvent, error)
}

type ShipmentEvents struct {
	Repo repository.ShipmentEventRepository
}

func NewShipmentEventsRepository(repo repository.ShipmentEventRepository) IShipmentEventsUseCase {
	return &ShipmentEvents{
		Repo: repo,
	}
}

func (uc *ShipmentEvents) GetByShipmentID(ctx context.Context, ShipmentEventID string) ([]entity.ShipmentEvent, error) {
	shipmentEvent, err := uc.Repo.GetByShipmentID(ctx, ShipmentEventID)
	if err != nil {
		return nil, err
	}

	return shipmentEvent, nil
}
