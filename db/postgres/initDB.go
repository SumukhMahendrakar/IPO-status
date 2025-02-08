package postgres

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectTODB() (*gorm.DB, error) {
	logrus.Info("Connecting to DB...")

	err := godotenv.Load()
	if err != nil {
		logrus.Errorln("Error loading env", err)
		return nil, err
	}

	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Errorln("Error connecting to DB", err)
		return nil, err
	}

	logrus.Infoln("Connected to DB successfully")
	return DB, nil
}
