package admin

type AdminRepository interface {
}

func NewAdminRepo() AdminRepository {
	return &adminRepo{}
}

type adminRepo struct {
}
