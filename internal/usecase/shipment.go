package usecase

import (
	"context"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
	"github.com/RajendraArkara/shipment_tracking/internal/repository"
)

type IShipmentUseCase interface {
	GetAll(ctx context.Context) ([]entity.Shipment, error)
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
