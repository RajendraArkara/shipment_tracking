package usecase

import (
	"context"

	"github.com/RajendraArkara/shipment_tracking/internal/entity"
	"github.com/RajendraArkara/shipment_tracking/internal/repository"
)

type IWebhookUseCase interface {
	Create(ctx context.Context, data *entity.Webhook) (*entity.Webhook, error)
}

type WebhookUseCase struct {
	Repo repository.WebhookSubscriptionsRepository
}

func NewWebhookRepository(repo repository.WebhookSubscriptionsRepository) IWebhookUseCase {
	return &WebhookUseCase{
		Repo: repo,
	}
}

func (uc *WebhookUseCase) Create(ctx context.Context, data *entity.Webhook) (*entity.Webhook, error) {
	webhook, err := uc.Repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return webhook, nil
}
