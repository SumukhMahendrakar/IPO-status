package dto

import (
	"github.com/SumukhMahendrakar/IPO-status/internal/dao"
	guuid "github.com/google/uuid"
)

type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	ID          guuid.UUID          `json:"id"`
	Name        string              `json:"name"`
	Email       string              `json:"email"`
	PhoneNumber string              `json:"phone_number"`
	PanNumbers  dao.PanNumbersArray `json:"pan_numbers"`
}

type IpoStatusReq struct {
	IpoName   string `json:"ipo_name"`
	PanNumber string `json:"pan_number"`
}

type IpoStatusResp struct {
	IpoName           string `json:"ipo_name"`
	PanNumber         string `json:"pan_number"`
	IsApplied         bool   `json:"is_applied"`
	IsAlloted         bool   `json:"is_alloted"`
	SecuritiesAlloted string `json:"securities_alloted"`
}
