package limit

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/limits", h.CreateLimit)
	rg.GET("/limits", h.GetLimits)
}

type CreateLimitRequest struct {
	TenorMonth int     `json:"tenor_month" binding:"required"`
	TotalLimit float64 `json:"total_limit" binding:"required"`
}

func (h *Handler) CreateLimit(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var req CreateLimitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateOrUpdateLimit(userID.(string), req.TenorMonth, req.TotalLimit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "limit berhasil disimpan"})
}

func (h *Handler) GetLimits(c *gin.Context) {
	userID, _ := c.Get("user_id")
	limits, err := h.service.GetAll(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, limits)
}
