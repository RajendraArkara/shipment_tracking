package main

import (
	"github.com/RajendraArkara/shipment_tracking/handler"
	shipmentEventsHttp "github.com/RajendraArkara/shipment_tracking/handler/v1/shipment_events/http"
	shipmentsHttp "github.com/RajendraArkara/shipment_tracking/handler/v1/shipments/http"
	"github.com/RajendraArkara/shipment_tracking/infrastructure/db"
	shipmentEventsRepo "github.com/RajendraArkara/shipment_tracking/infrastructure/repository/shipment_events"
	shipmentsRepo "github.com/RajendraArkara/shipment_tracking/infrastructure/repository/shipments"
	"github.com/RajendraArkara/shipment_tracking/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	shipmentEventsRepo := shipmentEventsRepo.NewShipmentEventsRepository(db.DB)

	shipmentRepo := shipmentsRepo.NewShipmentTrackingRepository(db.DB)
	shipmentUseCase := usecase.NewShipmentRepository(shipmentRepo, shipmentEventsRepo)
	shipmentHandler := shipmentsHttp.NewHandler(shipmentUseCase)

	shipmentEventsUseCase := usecase.NewShipmentEventsRepository(shipmentEventsRepo)
	shipmentEventsHandler := shipmentEventsHttp.NewHandler(shipmentEventsUseCase)

	server := gin.Default()
	handler.Shipments(server, shipmentHandler)
	handler.ShipmentEvents(server, shipmentEventsHandler)

	server.Run(":8080")
}
