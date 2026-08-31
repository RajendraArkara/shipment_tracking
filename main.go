package main

import (
	"github.com/RajendraArkara/shipment_tracking/handler"
	shipmentEventsHttp "github.com/RajendraArkara/shipment_tracking/handler/v1/shipment_events/http"
	shipmentsHttp "github.com/RajendraArkara/shipment_tracking/handler/v1/shipments/http"
	webhookSubscriptionsHttp "github.com/RajendraArkara/shipment_tracking/handler/v1/webhook_subscriptions/http"
	"github.com/RajendraArkara/shipment_tracking/infrastructure/db"
	shipmentEventsRepo "github.com/RajendraArkara/shipment_tracking/infrastructure/repository/shipment_events"
	shipmentsRepo "github.com/RajendraArkara/shipment_tracking/infrastructure/repository/shipments"
	webhookSubscriptionsRepo "github.com/RajendraArkara/shipment_tracking/infrastructure/repository/webhook_subscriptions"
	"github.com/RajendraArkara/shipment_tracking/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	shipmentEventsRepo := shipmentEventsRepo.NewShipmentEventsRepository(db.DB)
	webhookSubscriptionsRepo := webhookSubscriptionsRepo.NewWebhookRepository(db.DB)
	shipmentRepo := shipmentsRepo.NewShipmentTrackingRepository(db.DB)

	webhookSubscriptionsUseCase := usecase.NewWebhookRepository(webhookSubscriptionsRepo)
	webhookHandler := webhookSubscriptionsHttp.NewHandler(webhookSubscriptionsUseCase)

	shipmentUseCase := usecase.NewShipmentRepository(shipmentRepo, shipmentEventsRepo, webhookSubscriptionsRepo)
	shipmentHandler := shipmentsHttp.NewHandler(shipmentUseCase)

	shipmentEventsUseCase := usecase.NewShipmentEventsRepository(shipmentEventsRepo)
	shipmentEventsHandler := shipmentEventsHttp.NewHandler(shipmentEventsUseCase)

	server := gin.Default()
	handler.Shipments(server, shipmentHandler)
	handler.ShipmentEvents(server, shipmentEventsHandler)
	handler.WebhookSubscriptions(server, webhookHandler)

	server.Run(":8080")
}
