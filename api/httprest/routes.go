package httprest

import (
	"os"
	"time"

	"github.com/SumukhMahendrakar/IPO-status/api/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitRoutes() {
	logrus.Infoln("Initializing endpoints")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Cache CORS preflight response
	}))

	r.GET("/healthz", controllers.Health)
	r.POST("/login", controllers.UserLogin)
	r.POST("/get-ipo-status", controllers.GetIpoStatusController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	err := r.Run(":" + port)
	if err != nil {
		logrus.Errorln("Error on initializing endpoints, Err: ", err.Error())
		return
	}
}
