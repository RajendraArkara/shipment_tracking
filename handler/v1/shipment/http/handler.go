package shipmentTracking

import (
	"net/http"

	"github.com/RajendraArkara/shipment_tracking/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ShipmentHandler struct {
	UC usecase.IShipmentUseCase
}

func NewHandler(uc usecase.IShipmentUseCase) *ShipmentHandler {
	return &ShipmentHandler{
		UC: uc,
	}
}

func (h *ShipmentHandler) GetAll(ctx *gin.Context) {
	ship, err := h.UC.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "can not fetch the data",
			"error":   err.Error(),
		})
		return
	}

	resp := make([]ShipmentObject, 0, len(ship))

	for _, item := range ship {
		resp = append(resp, ShipmentObject{}.ParseFromEntity(item))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}
