package webhookSubscriptions

import (
	"context"
	"database/sql"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
	"github.com/RajendraArkara/shipment_tracking/internal/repository"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewWebhookRepository(db *sql.DB) repository.WebhookSubscriptionsRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(ctx context.Context, data *entity.Webhook) (*entity.Webhook, error) {
	query := `
		INSERT INTO webhook_subscriptions(shipment_id, target_url, secret_key)
		VALUES ($1, $2, $3)
		RETURNING id, shipment_id, target_url, secret_key, is_active, created_at
	`

	err := r.db.QueryRowContext(ctx, query, &data.ShipmentID, &data.TargetUrl, &data.SecretKey).Scan(
		&data.ID,
		&data.ShipmentID,
		&data.TargetUrl,
		&data.SecretKey,
		&data.IsActive,
		&data.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (r *PostgresRepository) GetByShipmentID(ctx context.Context, shipmentID string) ([]entity.Webhook, error) {
	query := `
		SELECT id, shipment_id, target_url, secret_key, is_active, created_at
		FROM webhook_subscriptions
		WHERE shipment_id = $1 AND is_active = true
	`

	rows, err := r.db.QueryContext(ctx, query, shipmentID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var webhooks []entity.Webhook

	for rows.Next() {
		var webhook entity.Webhook

		err = rows.Scan(
			&webhook.ID,
			&webhook.ShipmentID,
			&webhook.TargetUrl,
			&webhook.SecretKey,
			&webhook.IsActive,
			&webhook.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		webhooks = append(webhooks, webhook)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}
