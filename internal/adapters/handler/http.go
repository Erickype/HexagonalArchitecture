package handler

import "github.com/Erickype/HexagonalArchitecture/internal/core/services"

type HttpHandler struct {
	service services.MessengerService
}

func NewHttpHandler(service services.MessengerService) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}
