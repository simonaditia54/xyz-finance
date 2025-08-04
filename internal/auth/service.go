package auth

import (
	"errors"

	"xyz-finance/internal/model"
)

type Repository interface {
	FindByNIK(nik string) (*model.Consumer, error)
}

type Service interface {
	Login(nik string) (string, error)
}

type service struct {
	repo      Repository
	jwtSecret string
}

func NewService(repo Repository, secret string) Service {
	return &service{
		repo:      repo,
		jwtSecret: secret,
	}
}

func (s *service) Login(nik string) (string, error) {
	user, err := s.repo.FindByNIK(nik)
	if err != nil {
		return "", errors.New("invalid NIK")
	}

	token, err := GenerateJWT(user.ID, s.jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}
