package consumer

import (
	"time"

	"xyz-finance/internal/model"

	"github.com/google/uuid"
)

type Service interface {
	CreateConsumer(input model.Consumer) (string, error)
	FindByNIK(nik string) (*model.Consumer, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateConsumer(input model.Consumer) (string, error) {
	input.ID = uuid.NewString()
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	if err := s.repo.Save(input); err != nil {
		return "", err
	}
	return input.ID, nil
}

func (s *service) FindByNIK(nik string) (*model.Consumer, error) {
	return s.repo.FindByNIK(nik)
}
