package handler

import (
	shipmentsHttp "github.com/RajendraArkara/shipment_tracking/handler/v1/shipment/http"
	"github.com/gin-gonic/gin"
)

func Shipments(server *gin.Engine, h *shipmentsHttp.ShipmentHandler) {
	//shipments

	server.GET("/shipments/all", h.GetAll)
}
