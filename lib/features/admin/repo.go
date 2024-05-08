package admin


type AdminRepo interface {
}

func NewAdminRepo() AdminRepo {
	return &adminRepo{}
}

type adminRepo struct {
}