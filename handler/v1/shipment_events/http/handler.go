package shipmentevent

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/RajendraArkara/shipment_tracking/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ShipmentEventsHandler struct {
	UC usecase.IShipmentEventsUseCase
}

func NewHandler(uc usecase.IShipmentEventsUseCase) *ShipmentEventsHandler {
	return &ShipmentEventsHandler{
		UC: uc,
	}
}

func (h *ShipmentEventsHandler) GetByShipmentID(ctx *gin.Context) {
	id := ctx.Param("id")

	shipmentEvents, err := h.UC.GetByShipmentID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "id not found",
				"error":   err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "can not fetch the data",
			"error":   err.Error(),
		})
		return
	}

	resp := make([]ShipmentEventObject, 0, len(shipmentEvents))

	for _, item := range shipmentEvents {
		resp = append(resp, ShipmentEventObject{}.ParseFromEntity(item))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"shipment_event": resp,
	})
}
