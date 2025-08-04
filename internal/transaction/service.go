// internal/transaction/service.go
package transaction

import (
	"errors"
	"sync"

	"xyz-finance/internal/model"
)

type Service interface {
	CreateTransaction(tx model.Transaction) error
	GetTransactionsByUser(consumerID string) []model.Transaction
}

type service struct {
	repo Repository
	mu   sync.Mutex // for concurrency-safe limit update
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTransaction(tx model.Transaction) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Ambil limit berdasarkan consumer + tenor
	limit, err := s.repo.GetLimit(tx.ConsumerID, tx.TenorMonth)
	if err != nil {
		return err
	}

	// Hitung total yang dibutuhkan
	total := tx.JumlahOTR + tx.AdminFee + tx.JumlahBunga

	// Validasi limit
	if limit.TotalLimit-limit.UsedLimit < total {
		return errors.New("limit tidak mencukupi")
	}

	// Simpan transaksi
	if err := s.repo.SaveTransaction(tx); err != nil {
		return err
	}

	// Update used_limit
	limit.UsedLimit += total
	return s.repo.UpdateLimit(limit)
}

func (s *service) GetTransactionsByUser(consumerID string) []model.Transaction {
	var result []model.Transaction
	for _, tx := range s.repo.FindAll() {
		if tx.ConsumerID == consumerID {
			result = append(result, tx)
		}
	}
	return result
}
