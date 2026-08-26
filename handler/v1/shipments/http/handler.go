package shipmentTracking

import (
	"database/sql"
	"errors"
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

func (h *ShipmentHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	shipment, err := h.UC.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "id not found",
				"err":     err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "can not fetch the data",
			"error":   err.Error(),
		})
		return
	}

	resp := ShipmentObject{}.ParseFromEntity(*shipment)

	ctx.JSON(http.StatusOK, gin.H{
		"shipment": resp,
	})
}

func (h *ShipmentHandler) Create(ctx *gin.Context) {
	var req CreateShipmentRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	shipment := req.ToEntity()

	createShipment, err := h.UC.Create(ctx.Request.Context(), shipment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not create shipment",
			"err":     err.Error(),
		})
		return
	}

	parse := ShipmentObject{}.ParseFromEntity(*createShipment)

	ctx.JSON(http.StatusCreated, gin.H{
		"shipment": parse,
	})
}
