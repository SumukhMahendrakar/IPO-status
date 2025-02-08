package services

import "github.com/SumukhMahendrakar/IPO-status/internal/dto"

type UserServices interface {
	UserLogin(req *dto.UserLoginReq) (*dto.UserLoginResp, error, bool)
	IpoStatusCheck(req *dto.IpoStatusReq) (*dto.IpoStatusResp, error)
}
