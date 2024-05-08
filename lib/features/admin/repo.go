package admin

import "standup-api/lib/common/database"

type AdminRepository interface {
}

func NewAdminRepo(db *database.Database) AdminRepository {
	return &adminRepo{
		db: db,
	}
}

type adminRepo struct {
	db *database.Database
}
