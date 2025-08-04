package consumer

import (
	"errors"

	"xyz-finance/internal/model"

	"gorm.io/gorm"
)

type Repository interface {
	Save(c model.Consumer) error
	FindByID(id string) (*model.Consumer, error)
	FindByNIK(nik string) (*model.Consumer, error)
}

type mysqlRepo struct {
	db *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) Repository {
	db.AutoMigrate(&model.Consumer{}) // Auto migrate ke table
	return &mysqlRepo{db: db}
}

func (r *mysqlRepo) Save(c model.Consumer) error {
	return r.db.Create(&c).Error
}

func (r *mysqlRepo) FindByID(id string) (*model.Consumer, error) {
	var c model.Consumer
	err := r.db.First(&c, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *mysqlRepo) FindByNIK(nik string) (*model.Consumer, error) {
	var c model.Consumer
	err := r.db.First(&c, "nik = ?", nik).Error
	if err != nil {
		return nil, errors.New("not found")
	}
	return &c, nil
}
