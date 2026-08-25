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
}

type ShipmentUseCase struct {
	Repo repository.ShipmentRepository
}

func NewShipmentRepository(repo repository.ShipmentRepository) IShipmentUseCase {
	return &ShipmentUseCase{
		Repo: repo,
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
