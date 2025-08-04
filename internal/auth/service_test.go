package auth

import (
	"errors"
	"testing"
	"time"

	"xyz-finance/internal/model"

	"github.com/stretchr/testify/assert"
)

type fakeRepo struct {
	nikToConsumer map[string]model.Consumer
}

func (f *fakeRepo) FindByNIK(nik string) (*model.Consumer, error) {
	c, ok := f.nikToConsumer[nik]
	if !ok {
		return nil, errors.New("not found")
	}
	return &c, nil
}

func TestLogin_Berhasil(t *testing.T) {
	repo := &fakeRepo{
		nikToConsumer: map[string]model.Consumer{
			"123456": {
				ID:        "user-123",
				NIK:       "123456",
				FullName:  "Budi",
				CreatedAt: time.Now(),
			},
		},
	}

	svc := NewService(repo, "testsecret")
	token, err := svc.Login("123456")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestLogin_GagalNIKTidakDitemukan(t *testing.T) {
	repo := &fakeRepo{
		nikToConsumer: map[string]model.Consumer{},
	}

	svc := NewService(repo, "testsecret")
	token, err := svc.Login("999999")

	assert.Error(t, err)
	assert.Equal(t, "", token)
	assert.EqualError(t, err, "invalid NIK")
}
