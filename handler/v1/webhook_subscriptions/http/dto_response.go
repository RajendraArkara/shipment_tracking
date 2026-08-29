package webhookSubscriptions

import (
	"time"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
)

type WebhookObject struct {
	ID         string    `json:"webhook_subscription_id"`
	ShipmentID string    `json:"shipment_id"`
	TargetUrl  string    `json:"target_url"`
	SecretKey  string    `json:"secret_key"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
}

func (WebhookObject) ParseFromEntity(e entity.Webhook) WebhookObject {
	return WebhookObject{
		ID:         e.ID,
		ShipmentID: e.ShipmentID,
		TargetUrl:  e.TargetUrl,
		SecretKey:  e.SecretKey,
		IsActive:   e.IsActive,
		CreatedAt:  e.CreatedAt,
	}
}
