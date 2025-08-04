package transaction

import (
	"errors"
	"testing"
	"time"

	"xyz-finance/internal/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type fakeRepo struct {
	limit        *model.Limit
	transactions []model.Transaction
}

func (f *fakeRepo) GetLimit(consumerID string, tenor int) (*model.Limit, error) {
	if f.limit == nil {
		return nil, errors.New("not found")
	}
	return f.limit, nil
}

func (f *fakeRepo) SaveTransaction(tx model.Transaction) error {
	f.transactions = append(f.transactions, tx)
	return nil
}

func (f *fakeRepo) UpdateLimit(limit *model.Limit) error {
	f.limit = limit
	return nil
}

func (f *fakeRepo) FindAllByConsumer(consumerID string) ([]model.Transaction, error) {
	return f.transactions, nil
}

func (f *fakeRepo) FindAll() []model.Transaction {
	return f.transactions
}

func TestCreateTransaction_Success(t *testing.T) {
	repo := &fakeRepo{
		limit: &model.Limit{
			ID:         uuid.NewString(),
			ConsumerID: "user123",
			TenorMonth: 3,
			TotalLimit: 2000000,
			UsedLimit:  500000,
		},
	}

	service := NewService(repo)

	tx := model.Transaction{
		ID:            uuid.NewString(),
		ConsumerID:    "user123",
		TenorMonth:    3,
		JumlahOTR:     1000000,
		AdminFee:      50000,
		JumlahCicilan: 300000,
		JumlahBunga:   10000,
		NamaAsset:     "Motor",
		CreatedAt:     time.Now(),
	}

	err := service.CreateTransaction(tx)

	assert.NoError(t, err)
	// assert.Equal(t, 1500000.0, repo.limit.UsedLimit)
	assert.Equal(t, 1560000.0, repo.limit.UsedLimit)
	assert.Len(t, repo.transactions, 1)
}

func TestCreateTransaction_LimitTidakCukup(t *testing.T) {
	repo := &fakeRepo{
		limit: &model.Limit{
			ConsumerID: "user123",
			TenorMonth: 3,
			TotalLimit: 1000000,
			UsedLimit:  900000,
		},
	}

	service := NewService(repo)

	tx := model.Transaction{
		ID:            uuid.NewString(),
		ConsumerID:    "user123",
		TenorMonth:    3,
		JumlahOTR:     500000,
		AdminFee:      50000,
		JumlahCicilan: 300000,
		JumlahBunga:   10000,
		NamaAsset:     "Motor",
		CreatedAt:     time.Now(),
	}

	err := service.CreateTransaction(tx)

	assert.Error(t, err)
	assert.EqualError(t, err, "limit tidak mencukupi")

}
