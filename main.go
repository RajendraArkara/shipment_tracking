package main

import (
	"github.com/RajendraArkara/shipment_tracking/handler"
	shipmentsHttp "github.com/RajendraArkara/shipment_tracking/handler/v1/shipment/http"
	"github.com/RajendraArkara/shipment_tracking/infrastructure/db"
	shipmentsRepo "github.com/RajendraArkara/shipment_tracking/infrastructure/repository/shipment"
	"github.com/RajendraArkara/shipment_tracking/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	shipmentRepo := shipmentsRepo.NewShipmentTrackingRepository(db.DB)
	shipmentUseCase := usecase.NewShipmentRepository(shipmentRepo)
	shipmentHandler := shipmentsHttp.NewHandler(shipmentUseCase)

	server := gin.Default()
	handler.Shipments(server, shipmentHandler)

	server.Run(":8080")
}
