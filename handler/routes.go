package handler

import (
	shipmentEventsHttp "github.com/RajendraArkara/shipment_tracking/handler/v1/shipment_events/http"
	shipmentsHttp "github.com/RajendraArkara/shipment_tracking/handler/v1/shipments/http"
	"github.com/gin-gonic/gin"
)

func Shipments(server *gin.Engine, h *shipmentsHttp.ShipmentHandler) {
	//shipments
	server.GET("/shipments/all", h.GetAll)
	server.GET("/shipments/:id", h.GetByID)
	server.POST("/shipments", h.Create)
	server.PATCH("/shipments/:id/status", h.UpdateStatus)
}

func ShipmentEvents(server *gin.Engine, h *shipmentEventsHttp.ShipmentEventsHandler) {
	//shipmentEvents
	server.GET("/shipments/:id/history", h.GetByShipmentID)

}
