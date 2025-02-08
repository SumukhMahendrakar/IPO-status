package controllers

import (
	"net/http"

	"github.com/SumukhMahendrakar/IPO-status/initconf"
	"github.com/SumukhMahendrakar/IPO-status/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "UP"})
}

func UserLogin(c *gin.Context) {
	logrus.Infoln("Request received to login")
	var userLoginReq dto.UserLoginReq
	err := c.BindJSON(&userLoginReq)
	if err != nil {
		logrus.Errorln("Error reading the payload")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid Payload",
		})
		return
	}

	logrus.Infoln("The request payload binded successfully", userLoginReq)
	resp, err, pass := initconf.UsecaseContainer.UserService.UserLogin(&userLoginReq)
	if err != nil {
		logrus.Errorln("Error getting response from service", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Internal Error",
		})
		return
	}
	if !pass {
		logrus.Warnln("Password did not match")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "wrong password",
		})
		return
	}

	logrus.Infoln("successfully recieved the response from service", resp)
	c.JSON(http.StatusAccepted, gin.H{
		"data": resp,
	})
}

func GetIpoStatusController(c *gin.Context) {
	logrus.Infoln("Received request to get ipo status")
	var statusReq *dto.IpoStatusReq
	err := c.BindJSON(&statusReq)
	if err != nil {
		logrus.Errorln("Error reading the payload")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid Payload",
		})
		return
	}

	logrus.Infoln("Successfully binded json data", statusReq)

	resp, err := initconf.UsecaseContainer.UserService.IpoStatusCheck(statusReq)
	if err != nil {
		logrus.Errorln("Error getting response from service", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"data": resp,
	})

}
