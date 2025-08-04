package limit

import (
	"errors"

	"xyz-finance/internal/model"

	"gorm.io/gorm"
)

type Repository interface {
	Upsert(limit *model.Limit) error
	FindByConsumerAndTenor(consumerID string, tenor int) (*model.Limit, error)
	FindAllByConsumer(consumerID string) ([]model.Limit, error)
}

type mysqlRepo struct {
	db *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) Repository {
	return &mysqlRepo{db: db}
}

func (r *mysqlRepo) Upsert(limit *model.Limit) error {
	var existing model.Limit
	err := r.db.Where("consumer_id = ? AND tenor_month = ?", limit.ConsumerID, limit.TenorMonth).First(&existing).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return r.db.Create(limit).Error
	}
	if err != nil {
		return err
	}

	existing.TotalLimit = limit.TotalLimit
	existing.UpdatedAt = limit.UpdatedAt
	return r.db.Save(&existing).Error
}

func (r *mysqlRepo) FindByConsumerAndTenor(consumerID string, tenor int) (*model.Limit, error) {
	var limit model.Limit
	err := r.db.Where("consumer_id = ? AND tenor_month = ?", consumerID, tenor).First(&limit).Error
	if err != nil {
		return nil, err
	}
	return &limit, nil
}

func (r *mysqlRepo) FindAllByConsumer(consumerID string) ([]model.Limit, error) {
	var limits []model.Limit
	err := r.db.Where("consumer_id = ?", consumerID).Find(&limits).Error
	return limits, err
}
