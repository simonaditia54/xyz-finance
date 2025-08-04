package limit

import (
	"time"

	"xyz-finance/internal/model"

	"github.com/google/uuid"
)

type Service interface {
	CreateOrUpdateLimit(consumerID string, tenor int, total float64) error
	GetAll(consumerID string) ([]model.Limit, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateOrUpdateLimit(consumerID string, tenor int, total float64) error {
	limit := &model.Limit{
		ID:         uuid.NewString(),
		ConsumerID: consumerID,
		TenorMonth: tenor,
		TotalLimit: total,
		UsedLimit:  0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	return s.repo.Upsert(limit)
}

func (s *service) GetAll(consumerID string) ([]model.Limit, error) {
	return s.repo.FindAllByConsumer(consumerID)
}
