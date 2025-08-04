package transaction

import (
	"net/http"
	"time"

	"xyz-finance/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r gin.IRoutes) {
	r.POST("/transactions", h.CreateTransaction)
	r.GET("/transactions", h.GetTransactions)
}

type CreateTransactionRequest struct {
	// ConsumerID     string  `json:"consumer_id" binding:"required"`
	TenorMonth     int     `json:"tenor_month" binding:"required"`
	ContractNumber string  `json:"contract_number" binding:"required"`
	JumlahOTR      float64 `json:"jumlah_otr" binding:"required"`
	AdminFee       float64 `json:"admin_fee"`
	JumlahCicilan  float64 `json:"jumlah_cicilan"`
	JumlahBunga    float64 `json:"jumlah_bunga"`
	NamaAsset      string  `json:"nama_asset"`
}

func (h *Handler) CreateTransaction(c *gin.Context) {
	var req CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")

	tx := model.Transaction{
		ID:             uuid.NewString(),
		ContractNumber: req.ContractNumber,
		ConsumerID:     userID.(string), // req.ConsumerID,
		TenorMonth:     req.TenorMonth,
		JumlahOTR:      req.JumlahOTR,
		AdminFee:       req.AdminFee,
		JumlahCicilan:  req.JumlahCicilan,
		JumlahBunga:    req.JumlahBunga,
		NamaAsset:      req.NamaAsset,
		CreatedAt:      time.Now(),
	}

	if err := h.service.CreateTransaction(tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "transaksi berhasil disimpan"})
}

func (h *Handler) GetTransactions(c *gin.Context) {
	userID, _ := c.Get("user_id")
	list := h.service.GetTransactionsByUser(userID.(string))

	c.JSON(http.StatusOK, gin.H{
		"data": list,
	})
}
