package repository

import (
	"context"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
)

type WebhookSubscriptionsRepository interface {
	Create(ctx context.Context, data *entity.Webhook) (*entity.Webhook, error)
}
