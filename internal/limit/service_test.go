package limit

import (
	"testing"

	"xyz-finance/internal/model"

	"github.com/stretchr/testify/assert"
)

type fakeRepo struct {
	UpsertCalled     bool
	UpsertReceived   *model.Limit
	FindAllCalled    bool
	ConsumerIDPassed string
	Limit            []model.Limit
	Err              error
}

func (f *fakeRepo) Upsert(limit *model.Limit) error {
	f.UpsertCalled = true
	f.UpsertReceived = limit
	return f.Err
}

func (f *fakeRepo) FindByConsumerAndTenor(_ string, _ int) (*model.Limit, error) {
	return nil, nil
}

func (f *fakeRepo) FindAllByConsumer(consumerID string) ([]model.Limit, error) {
	f.FindAllCalled = true
	f.ConsumerIDPassed = consumerID
	return f.Limit, f.Err
}

func TestCreateOrUpdateLimit(t *testing.T) {
	mock := &fakeRepo{}
	svc := NewService(mock)

	err := svc.CreateOrUpdateLimit("user-1", 6, 3000000)

	assert.NoError(t, err)
	assert.True(t, mock.UpsertCalled)
	assert.Equal(t, "user-1", mock.UpsertReceived.ConsumerID)
	assert.Equal(t, 6, mock.UpsertReceived.TenorMonth)
	assert.Equal(t, 3000000.0, mock.UpsertReceived.TotalLimit)
}

func TestGetAll(t *testing.T) {
	mock := &fakeRepo{
		Limit: []model.Limit{
			{ID: "limit-1", ConsumerID: "user-1"},
		},
	}

	svc := NewService(mock)
	limits, err := svc.GetAll("user-1")

	assert.NoError(t, err)
	assert.True(t, mock.FindAllCalled)
	assert.Equal(t, "user-1", mock.ConsumerIDPassed)
	assert.Len(t, limits, 1)
}
