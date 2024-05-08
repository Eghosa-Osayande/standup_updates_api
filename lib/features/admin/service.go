package admin

import (
	"errors"
	"os"
	"standup-api/lib/utils/jwt"
	"time"
)

type AdminService interface {
	Login(*AdminLoginInputDto) (*AdminLoginOutputDto, error)
}

func NewAdminService(repo AdminRepo) AdminService {
	return &adminService{}
}

type adminService struct {
}

func (s *adminService) Login(input *AdminLoginInputDto) (*AdminLoginOutputDto, error) {
	adminPassword, ok := os.LookupEnv("ADMIN_PASSWORD")
	if !ok {
		// log
		return nil, errors.New("error generating token")
	}

	if input.Password != adminPassword {
		return nil, errors.New("invalid password")
	}

	secret, ok := os.LookupEnv("TOKEN_KEY")
	if !ok {
		// log
		return nil, errors.New("error generating token")
	}

	token, err := jwt.GenerateJwtToken(map[string]any{"role": "admin"}, secret, 60*time.Minute)

	if err != nil {
		return nil, errors.New("error generating token")
	}

	return &AdminLoginOutputDto{Token: token}, nil
}
