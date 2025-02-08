package utils

import (
	"fmt"
	"time"

	"github.com/SumukhMahendrakar/IPO-status/internal/dto"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/sirupsen/logrus"
)

func IpoStatusCheker(ipoName string, pan string) *dto.IpoStatusResp {
	logrus.Infoln("Checking the registrar for IPO")
	reg := IpoMapper[ipoName]
	switch reg {
	case "linkintime":
		res := LinkinTimeScraper(ipoName, pan)
		logrus.Infoln(res)
		return res
	case "kfintech":
		res := KFintechScraper(ipoName, pan)
		logrus.Infoln(res)
		return res
	default:
		logrus.Infoln("IPO not found")
		return nil
	}
}

func LinkinTimeScraper(ipoName string, pan string) *dto.IpoStatusResp {
	logrus.Infoln("Visiting Linkintime registrar")

	url := launcher.New().Leakless(false).Headless(false).MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()

	page := browser.MustPage("https://linkintime.co.in/initial_offer/public-issues.html").MustWaitStable()

	defer browser.MustClose()

	page.MustElement("#ddlCompany").MustClick().MustSelect(ipoName)
	logrus.Infoln("Found the IPO")

	page.MustElement("input#txtStat").MustClick().MustInput(pan)

	page.MustElement("input#btnsearc").MustClick()

	tablefound, _, _ := page.HasX("//table//*[contains(text(), 'Application Details - HNI')]")
	if !tablefound {
		logrus.Infoln("Record not found")
		return &dto.IpoStatusResp{
			IpoName:           ipoName,
			PanNumber:         pan,
			IsApplied:         false,
			IsAlloted:         false,
			SecuritiesAlloted: "0",
		}
	}

	tbody := page.MustElement("tbody")

	res := tbody.MustElement("tr:nth-of-type(4)").MustElement("td:nth-of-type(2)")

	if res.MustText() == "0" {
		return &dto.IpoStatusResp{
			IpoName:           ipoName,
			PanNumber:         pan,
			IsApplied:         true,
			IsAlloted:         false,
			SecuritiesAlloted: res.MustText(),
		}
	}

	return &dto.IpoStatusResp{
		IpoName:           ipoName,
		PanNumber:         pan,
		IsApplied:         true,
		IsAlloted:         true,
		SecuritiesAlloted: res.MustText(),
	}
}

func KFintechScraper(ipoName string, pan string) *dto.IpoStatusResp {
	logrus.Infoln("Visiting KFintech registrar")

	url := launcher.New().Leakless(false).Headless(false).MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	// defer browser.MustClose()

	page := browser.MustPage("https://ipostatus.kfintech.com/").MustWaitStable()

	page.MustElement("select#ddl_ipo").MustClick().MustSelect(ipoName)

	logrus.Infoln("Found the IPO")

	page.MustElement("input#pan").MustClick()

	page.MustElement("input#txt_pan").MustInput(pan)

	var isValidCaptcha bool = false

	for !isValidCaptcha {
		page.MustElement("img#captchaimg").MustScreenshot("captcha.png")

		resp, err := PerformOCR()
		if err != nil {
			fmt.Println("error performing ocr", err.Error())
			return nil
		}

		fmt.Println("OCR Succesfull", resp)

		page.MustElement("input#txt_captcha").MustInput(resp)

		page.MustElement("a#btn_submit_query").MustClick()

		time.Sleep(2 * time.Second)

		invalidCaptcha, _, _ := page.HasX("//div//*[contains(translate(text(), 'abcdefghijklmnopqrstuvwxyz', 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'), 'CAPTCHA IS INVALID') or contains(translate(text(), 'abcdefghijklmnopqrstuvwxyz', 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'), 'PLEASE ENTER CAPTCHA') or contains(translate(text(), 'abcdefghijklmnopqrstuvwxyz', 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'), 'CAPTCHA IS INVALID OR EXPIRED')]")

		if invalidCaptcha {
			fmt.Println("captcha was invalid")

			okButton, err := page.ElementX("//button[translate(text(), 'abcdefghijklmnopqrstuvwxyz', 'ABCDEFGHIJKLMNOPQRSTUVWXYZ')='OK']")
			if err == nil {
				fmt.Println("OK button detected")
				okButton.MustClick()
			}

			fmt.Println("OK button not found")

			continue
		}

		isValidCaptcha = true
	}

	noPanDetails, _, _ := page.HasX("//div//*[contains(translate(text(), 'abcdefghijklmnopqrstuvwxyz', 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'), 'PAN DETAILS  NOT AVAILABLE.')]")

	if noPanDetails {
		fmt.Println("Pan details are not available")
		return &dto.IpoStatusResp{
			IpoName:           ipoName,
			PanNumber:         pan,
			IsApplied:         false,
			IsAlloted:         false,
			SecuritiesAlloted: "0",
		}
	}

	securitiesAlloted := page.MustElement("span#grid_results_lbl_allot_0")

	logrus.Infoln("share alloted is : ", securitiesAlloted.MustText())

	if securitiesAlloted.MustText() == "0" {
		return &dto.IpoStatusResp{
			IpoName:           ipoName,
			PanNumber:         pan,
			IsApplied:         true,
			IsAlloted:         false,
			SecuritiesAlloted: "0",
		}
	}
	return &dto.IpoStatusResp{
		IpoName:           ipoName,
		PanNumber:         pan,
		IsApplied:         true,
		IsAlloted:         true,
		SecuritiesAlloted: securitiesAlloted.MustText(),
	}
}
