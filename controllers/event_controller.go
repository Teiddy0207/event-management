package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"be-event/services"
)

type EventController struct {
	service *services.EventService
}

func NewEventController(service *services.EventService) *EventController {
	return &EventController{service}
}

func (c *EventController) ListEvents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "ListEvents dummy"})
}
