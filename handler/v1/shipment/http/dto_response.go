package shipmentTracking

import (
	"time"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
)

type ShipmentObject struct {
	ShipmentID    string     `json:"shipment_id"`
	OrderID       string     `json:"order_id"`
	CarrierID     string     `json:"carrier_id"`
	Origin        string     `json:"origin"`
	Destination   string     `json:"destination"`
	CurrentStatus string     `json:"current_status"`
	Eta           *time.Time `json:"eta"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func (ShipmentObject) ParseFromEntity(e entity.Shipment) ShipmentObject {
	return ShipmentObject{
		ShipmentID:    e.ShipmentID,
		OrderID:       e.OrderID,
		CarrierID:     e.CarrierID,
		Origin:        e.Origin,
		Destination:   e.Destination,
		CurrentStatus: e.CurrentStatus,
		Eta:           e.Eta,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
	}
}
