package consumer

import (
	"net/http"

	"xyz-finance/internal/model"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r gin.IRoutes) {
	r.POST("/consumers", h.CreateConsumer)
	r.GET("/consumers/:nik", h.GetByNIK)
}

func (h *Handler) CreateConsumer(c *gin.Context) {
	var input model.Consumer

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.CreateConsumer(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "consumer created", "id": id})
}

func (h *Handler) GetByNIK(c *gin.Context) {
	nik := c.Param("nik")

	consumer, err := h.service.FindByNIK(nik)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "consumer not found"})
		return
	}

	c.JSON(http.StatusOK, consumer)
}
