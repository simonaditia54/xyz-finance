package transaction

import (
	"xyz-finance/internal/model"

	"gorm.io/gorm"
)

type Repository interface {
	GetLimit(consumerID string, tenor int) (*model.Limit, error)
	UpdateLimit(limit *model.Limit) error
	SaveTransaction(tx model.Transaction) error
	FindAll() []model.Transaction
}

type mysqlRepo struct {
	db *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) Repository {
	db.AutoMigrate(&model.Transaction{}, &model.Limit{}) // auto create tables
	return &mysqlRepo{db: db}
}

func (r *mysqlRepo) SaveTransaction(tx model.Transaction) error {
	return r.db.Create(&tx).Error
}

func (r *mysqlRepo) FindAll() []model.Transaction {
	var txs []model.Transaction
	r.db.Find(&txs)
	return txs
}

func (r *mysqlRepo) GetLimit(consumerID string, tenor int) (*model.Limit, error) {
	var limit model.Limit
	err := r.db.Where("consumer_id = ? AND tenor_month = ?", consumerID, tenor).First(&limit).Error
	if err != nil {
		return nil, err
	}
	return &limit, nil
}

func (r *mysqlRepo) UpdateLimit(limit *model.Limit) error {
	return r.db.Save(limit).Error
}
