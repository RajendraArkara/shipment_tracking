package webhookSubscriptions

import "github.com/RajendraArkara/shipment_tracking/internal/entity"

type WebhookRequest struct {
	ShipmentID string `json:"shipment_id" binding:"required"`
	TargetUrl  string `json:"target_url" binding:"required"`
	SecretKey  string `json:"secret_key" binding:"required"`
}

func (dto WebhookRequest) ToEntity() *entity.Webhook {
	return &entity.Webhook{
		ShipmentID: dto.ShipmentID,
		TargetUrl:  dto.TargetUrl,
		SecretKey:  dto.SecretKey,
	}
}
