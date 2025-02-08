package main

import (
	"github.com/SumukhMahendrakar/IPO-status/api/httprest"
	"github.com/SumukhMahendrakar/IPO-status/db/postgres"
	"github.com/SumukhMahendrakar/IPO-status/initconf"
)

func init() {
	postgres.ConnectTODB()

	initconf.InitApp()

	httprest.InitRoutes()
}

func main() {
}
