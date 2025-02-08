package serviceimpl

import (
	"errors"

	"github.com/SumukhMahendrakar/IPO-status/internal/dto"
	"github.com/SumukhMahendrakar/IPO-status/internal/repository"
	"github.com/SumukhMahendrakar/IPO-status/internal/services"
	"github.com/SumukhMahendrakar/IPO-status/internal/services/utils"
	"github.com/sirupsen/logrus"
)

type UserServicesImpl struct {
	userRepo repository.UserRepo
}

func NewUserServiceImpl(userRepo repository.UserRepo) services.UserServices {
	return &UserServicesImpl{
		userRepo: userRepo,
	}
}

func (userServicesImpl *UserServicesImpl) UserLogin(req *dto.UserLoginReq) (*dto.UserLoginResp, error, bool) {
	logrus.Infoln("Login service triggered", req)
	userInfo, err := userServicesImpl.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		logrus.Errorln("Error getting the user info", err.Error())
		return nil, err, true
	}
	logrus.Infoln("User info received")

	if req.Password != userInfo.Password {
		logrus.Infoln("Password does not match", req.Password)
		return nil, nil, false
	}

	logrus.Infoln("Password matched successfully")

	userLoginResp := &dto.UserLoginResp{
		ID:          userInfo.ID,
		Name:        userInfo.Name,
		Email:       userInfo.Email,
		PhoneNumber: userInfo.PhoneNumber,
		PanNumbers:  userInfo.PanNumbers,
	}
	return userLoginResp, nil, true
}

func (userServicesImpl *UserServicesImpl) IpoStatusCheck(req *dto.IpoStatusReq) (*dto.IpoStatusResp, error) {
	logrus.Infoln("Ipo Checker service called", req.IpoName)
	resp := utils.IpoStatusCheker(req.IpoName, req.PanNumber)
	if resp != nil {
		return resp, nil
	}

	return nil, errors.New("IPO not found")
}
