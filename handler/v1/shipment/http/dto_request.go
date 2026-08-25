package shipmentTracking

import (
	"time"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
)

type CreateShipmentRequest struct {
	OrderID     *string    `json:"order_id"`
	CarrierID   *string    `json:"carrier_id"`
	Origin      string     `json:"origin" binding:"required"`
	Destination string     `json:"destination" binding:"required"`
	Eta         *time.Time `json:"eta"`
}

func (dto CreateShipmentRequest) ToEntity() *entity.Shipment {
	return &entity.Shipment{
		OrderID:     dto.OrderID,
		CarrierID:   dto.CarrierID,
		Origin:      dto.Origin,
		Destination: dto.Destination,
		Eta:         dto.Eta,
	}
}
