package repoimpl

import (
	"github.com/SumukhMahendrakar/IPO-status/internal/dao"
	"github.com/SumukhMahendrakar/IPO-status/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const tableUsers string = "users"

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepoImpl(db *gorm.DB) repository.UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}

func (userRepo *UserRepoImpl) GetUserByEmail(email string) (*dao.User, error) {
	logrus.Infoln("Getting user from repo", email)
	var userInfo *dao.User
	query := "SELECT * FROM " + tableUsers + " WHERE email=?"
	res := userRepo.db.Raw(query, email).Scan(&userInfo)
	if res.Error != nil {
		logrus.Errorln("Error getting user info", res.Error.Error())
		return nil, res.Error
	}

	logrus.Infoln("Got user info from repo", userInfo)
	return userInfo, nil
}
