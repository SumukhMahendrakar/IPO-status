package repository

import "github.com/SumukhMahendrakar/IPO-status/internal/dao"

type UserRepo interface {
	GetUserByEmail(email string) (*dao.User, error)
}
