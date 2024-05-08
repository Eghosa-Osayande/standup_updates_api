package admin

import (
	"net/http"
	"os"
	"standup-api/lib/utils/http_response"
	"standup-api/lib/utils/jwt"
	"time"
)

type AdminService interface {
	Login(*AdminLoginInputDto) (*AdminLoginOutputDto, *http_response.HttpError)
}

func NewAdminService(repo AdminRepository) AdminService {
	return &adminService{}
}

type adminService struct {
}

func (s *adminService) Login(input *AdminLoginInputDto) (*AdminLoginOutputDto, *http_response.HttpError) {
	adminPassword, ok := os.LookupEnv("ADMIN_PASSWORD")
	if !ok {
		// log
		return nil, http_response.NewHttpError(http.StatusInternalServerError, "error generating token")
	}

	if input.Password != adminPassword {
		return nil, http_response.NewHttpError(http.StatusUnauthorized, "invalid password")
	}

	secret, ok := os.LookupEnv("TOKEN_KEY")
	if !ok {
		// log
		return nil, http_response.NewHttpError(http.StatusInternalServerError, "error generating token")
	}

	token, err := jwt.GenerateJwtToken(map[string]any{"role": "admin"}, secret, 60*time.Minute)

	if err != nil {
		return nil, http_response.NewHttpError(http.StatusInternalServerError, "error generating token")
	}

	return &AdminLoginOutputDto{Token: token}, nil
}
