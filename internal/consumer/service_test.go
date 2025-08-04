package consumer

import (
	"errors"
	"testing"

	"xyz-finance/internal/model"

	"github.com/stretchr/testify/assert"
)

type fakeRepo struct {
	Consumer *model.Consumer
	Err      error
}

func (f *fakeRepo) Save(c model.Consumer) error {
	f.Consumer = &c
	return f.Err
}

func (f *fakeRepo) FindByNIK(nik string) (*model.Consumer, error) {
	return nil, nil
}

func (f *fakeRepo) FindByID(id string) (*model.Consumer, error) {
	return nil, nil
}

func TestCreateConsumer_Berhasil(t *testing.T) {
	repo := &fakeRepo{}
	service := NewService(repo)

	input := model.Consumer{
		NIK:      "123456",
		FullName: "Budi",
	}

	id, err := service.CreateConsumer(input)

	assert.NoError(t, err)
	assert.NotEmpty(t, id)
	assert.NotNil(t, repo.Consumer)
	assert.Equal(t, "Budi", repo.Consumer.FullName)
	assert.False(t, repo.Consumer.CreatedAt.IsZero())
}

func TestCreateConsumer_RepoError(t *testing.T) {
	repo := &fakeRepo{
		Err: errors.New("database error"),
	}
	service := NewService(repo)

	input := model.Consumer{
		NIK:      "123456",
		FullName: "Budi",
	}

	id, err := service.CreateConsumer(input)

	assert.Error(t, err)
	assert.Equal(t, "", id)
	assert.EqualError(t, err, "database error")
}

func TestFindByNIK_Berhasil(t *testing.T) {
	repo := &fakeRepo{
		Consumer: &model.Consumer{
			ID:       "user-1",
			NIK:      "123456",
			FullName: "Budi Santoso",
		},
	}
	service := NewService(repo)

	result, err := service.FindByNIK("123456")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Budi Santoso", result.FullName)
}

func TestFindByNIK_TidakDitemukan(t *testing.T) {
	repo := &fakeRepo{
		Consumer: &model.Consumer{},
	}
	service := NewService(repo)

	result, err := service.FindByNIK("000000")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "not found")
}
