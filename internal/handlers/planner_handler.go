package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/immortal1405/vacay_planner/internal/models"
	"github.com/immortal1405/vacay_planner/internal/services/planner"
)

type Handler struct {
	plannerService *planner.Service
}

func NewHandler(plannerService *planner.Service) *Handler {
	return &Handler{
		plannerService: plannerService,
	}
}

func (h *Handler) GeneratePlan(c *gin.Context) {
	var req models.VacationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.plannerService.GeneratePlan(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetTravelTips(c *gin.Context) {
	var req models.TravelTipsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.plannerService.GetTravelTips(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) ChatWithPlanner(c *gin.Context) {
	var req models.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.plannerService.ChatWithPlanner(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
