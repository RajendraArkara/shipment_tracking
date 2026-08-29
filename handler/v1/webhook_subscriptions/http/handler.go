package webhookSubscriptions

import (
	"net/http"

	"github.com/RajendraArkara/shipment_tracking/internal/usecase"
	"github.com/gin-gonic/gin"
)

type WebHookHandler struct {
	Uc usecase.IWebhookUseCase
}

func NewHandler(uc usecase.IWebhookUseCase) *WebHookHandler {
	return &WebHookHandler{
		Uc: uc,
	}
}

func (h *WebHookHandler) Create(ctx *gin.Context) {
	var req WebhookRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	webhook := req.ToEntity()

	createWebhook, err := h.Uc.Create(ctx.Request.Context(), webhook)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "can not create webhook subscription",
			"error":   err.Error(),
		})
		return
	}

	parse := WebhookObject{}.ParseFromEntity(*createWebhook)

	ctx.JSON(http.StatusCreated, gin.H{
		"webhook_subs": parse,
	})
}
