package shipmentevents

import (
	"context"
	"database/sql"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
	"github.com/RajendraArkara/shipment_tracking/internal/repository"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewShipmentEventsRepository(db *sql.DB) repository.ShipmentEventRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) GetByShipmentID(ctx context.Context, ShipmentEventID string) ([]entity.ShipmentEvent, error) {
	query := `
		SELECT id, shipment_id, status, location, notes, occurred_at, created_at
		FROM shipment_events
		WHERE shipment_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query, ShipmentEventID)
	if err != nil {
		return nil, err
	}

	var shipmentevents []entity.ShipmentEvent

	defer rows.Close()

	for rows.Next() {
		var shipmentevent entity.ShipmentEvent

		err := rows.Scan(
			&shipmentevent.ShipmentEventID,
			&shipmentevent.ShipmentID,
			&shipmentevent.Status,
			&shipmentevent.Location,
			&shipmentevent.Notes,
			&shipmentevent.OccurredAt,
			&shipmentevent.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		shipmentevents = append(shipmentevents, shipmentevent)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return shipmentevents, nil
}

func (r *PostgresRepository) Create(ctx context.Context, data *entity.ShipmentEvent) error {
	query := `
		INSERT INTO shipment_events(shipment_id, status, location, notes)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.ExecContext(ctx, query, data.ShipmentID, data.Status, data.Location, data.Notes)
	return err
}
