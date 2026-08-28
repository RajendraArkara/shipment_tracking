package usecase

import (
	"context"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
	"github.com/RajendraArkara/shipment_tracking/internal/repository"
)

type IShipmentUseCase interface {
	GetAll(ctx context.Context) ([]entity.Shipment, error)
	GetByID(ctx context.Context, id string) (*entity.Shipment, error)
	Create(ctx context.Context, daata *entity.Shipment) (*entity.Shipment, error)
	UpdateStatus(ctx context.Context, shipments, newStatus string, location, notes *string) (*entity.Shipment, error)
}

type ShipmentUseCase struct {
	Repo              repository.ShipmentRepository
	ShipmentEventRepo repository.ShipmentEventRepository
}

func NewShipmentRepository(repo repository.ShipmentRepository, eventRepo repository.ShipmentEventRepository) IShipmentUseCase {
	return &ShipmentUseCase{
		Repo:              repo,
		ShipmentEventRepo: eventRepo,
	}
}

func (uc *ShipmentUseCase) GetAll(ctx context.Context) ([]entity.Shipment, error) {
	shipment, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}

func (uc *ShipmentUseCase) GetByID(ctx context.Context, id string) (*entity.Shipment, error) {
	shipment, err := uc.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}

func (uc *ShipmentUseCase) Create(ctx context.Context, data *entity.Shipment) (*entity.Shipment, error) {
	shipment, err := uc.Repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}

func (uc *ShipmentUseCase) UpdateStatus(ctx context.Context, shipmentID, newStatus string, location, notes *string) (*entity.Shipment, error) {
	shipment, err := uc.Repo.GetByID(ctx, shipmentID)
	if err != nil {
		return nil, err
	}

	if !shipment.CanTransitionTo(newStatus) {
		return nil, entity.ErrorInvalid
	}

	event := &entity.ShipmentEvent{
		ShipmentID: shipmentID,
		Status:     newStatus,
		Location:   location,
		Notes:      notes,
	}

	err = uc.ShipmentEventRepo.Create(ctx, event)
	if err != nil {
		return nil, err
	}

	updateShipment, err := uc.Repo.UpdateStatus(ctx, shipmentID, newStatus)
	if err != nil {
		return nil, err
	}

	return updateShipment, nil
}
