package shipmenttracking

import (
	"context"
	"database/sql"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
	"github.com/RajendraArkara/shipment_tracking/internal/repository"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewShipmentTrackingRepository(db *sql.DB) repository.ShipmentRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) GetAll(ctx context.Context) ([]entity.Shipment, error) {
	query := `
		SELECT id, order_id, carrier_id, origin, destination, current_status, eta, created_at, updated_at
		FROM shipments
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var shipments []entity.Shipment

	for rows.Next() {
		var shipment entity.Shipment

		err := rows.Scan(
			&shipment.ShipmentID,
			&shipment.OrderID,
			&shipment.CarrierID,
			&shipment.Origin,
			&shipment.Destination,
			&shipment.CurrentStatus,
			&shipment.Eta,
			&shipment.CreatedAt,
			&shipment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		shipments = append(shipments, shipment)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return shipments, nil
}
