package main

import (
	"github.com/RajendraArkara/shipment_tracking/handler"
	shipmentEventsHttp "github.com/RajendraArkara/shipment_tracking/handler/v1/shipment_events/http"
	shipmentsHttp "github.com/RajendraArkara/shipment_tracking/handler/v1/shipments/http"
	"github.com/RajendraArkara/shipment_tracking/infrastructure/db"
	shipmentsRepo "github.com/RajendraArkara/shipment_tracking/infrastructure/repository/shipment"
	shipmentEventsRepo "github.com/RajendraArkara/shipment_tracking/infrastructure/repository/shipment_events"
	"github.com/RajendraArkara/shipment_tracking/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	shipmentRepo := shipmentsRepo.NewShipmentTrackingRepository(db.DB)
	shipmentUseCase := usecase.NewShipmentRepository(shipmentRepo)
	shipmentHandler := shipmentsHttp.NewHandler(shipmentUseCase)

	shipmentEventsRepo := shipmentEventsRepo.NewShipmentEventsRepository(db.DB)
	shipmentEventsUseCase := usecase.NewShipmentEventsRepository(shipmentEventsRepo)
	shipmentEventsHandler := shipmentEventsHttp.NewHandler(shipmentEventsUseCase)

	server := gin.Default()
	handler.Shipments(server, shipmentHandler)
	handler.ShipmentEvents(server, shipmentEventsHandler)

	server.Run(":8080")
}
