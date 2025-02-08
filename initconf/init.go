package initconf

import (
	"github.com/SumukhMahendrakar/IPO-status/db/postgres"
	repoimpl "github.com/SumukhMahendrakar/IPO-status/internal/repository/repoImpl"
	"github.com/SumukhMahendrakar/IPO-status/internal/services"
	serviceimpl "github.com/SumukhMahendrakar/IPO-status/internal/services/serviceImpl"
	"github.com/sirupsen/logrus"
)

type ServiceContainer struct {
	UserService services.UserServices
}

var UsecaseContainer *ServiceContainer

func InitApp() {
	logrus.Infoln("Initialising the services")

	db, err := postgres.ConnectTODB()
	if err != nil {
		logrus.Error("error on initialization of DB")
		panic(err)
	}

	logrus.Infoln("Initialising repositories")

	// Repo Initialisaion
	userRepo := repoimpl.NewUserRepoImpl(db)

	// Usecase Initialisation
	userUsecase := serviceimpl.NewUserServiceImpl(userRepo)

	UsecaseContainer = &ServiceContainer{
		UserService: userUsecase,
	}

	logrus.Infoln("Initialised the app successfully")
}
