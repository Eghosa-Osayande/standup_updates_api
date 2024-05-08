package admin

import "errors"

type AdminService interface {
	Login(*AdminLoginInputDto) (*AdminLoginOutputDto, error)
}

func NewAdminService( repo AdminRepo) AdminService {
	return &adminService{}
}

type adminService struct {
	
}

func (s *adminService) Login(input *AdminLoginInputDto) (*AdminLoginOutputDto, error) {
	if input.Password != "1234" {

		return nil, errors.New("invalid password")
	}
	return &AdminLoginOutputDto{Token: ""}, nil
}
