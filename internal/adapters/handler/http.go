package handler

import (
	"github.com/Erickype/HexagonalArchitecture/internal/core/domain"
	"github.com/Erickype/HexagonalArchitecture/internal/core/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpHandler struct {
	service services.MessengerService
}

func (h *HttpHandler) SaveMessage(ctx *gin.Context) {
	var message domain.Message
	err := ctx.ShouldBindJSON(&message)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	err = h.service.SaveMessage(message)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"success": "message created",
	})
}

func (h *HttpHandler) ReadMessage(ctx *gin.Context) {
	id := ctx.Param("id")
	message, err := h.service.ReadMessage(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, message)
}

func NewHttpHandler(service services.MessengerService) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}
