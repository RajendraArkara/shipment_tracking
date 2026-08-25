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

func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*entity.Shipment, error) {
	query := `
		SELECT id, order_id, carrier_id, origin, destination, current_status, eta, created_at, updated_at
		FROM shipments
		WHERE id = $1
		ORDER BY created_at DESC
	`

	row := r.db.QueryRow(query, id)

	var shipment entity.Shipment

	err := row.Scan(
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

	return &shipment, nil
}

func (r *PostgresRepository) Create(ctx context.Context, data *entity.Shipment) (*entity.Shipment, error) {
	query := `
		INSERT INTO shipments (order_id, carrier_id, origin, destination, eta)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, order_id, carrier_id, origin, destination, current_status, eta, created_at, updated_at
	`

	err := r.db.QueryRow(
		query,
		data.OrderID,
		data.CarrierID,
		data.Origin,
		data.Destination,
		data.Eta,
	).Scan(&data.ShipmentID, &data.OrderID, &data.CarrierID, &data.Origin, &data.Destination, &data.CurrentStatus, &data.Eta, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return data, nil
}
