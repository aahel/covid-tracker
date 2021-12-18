package main

import (
	"github.com/aahel/covid-tracker/api"
	handler "github.com/aahel/covid-tracker/api/v1"
	"github.com/aahel/covid-tracker/config"
	"github.com/aahel/covid-tracker/server"
	"github.com/aahel/covid-tracker/service"
	"github.com/aahel/covid-tracker/store"
)

func main() {
	lgr := config.GetConsoleLogger()
	appConfig := config.GetAppConfig()
	dbConn := config.GetDBConn(lgr, appConfig.DB)
	router := config.InitRouter(lgr, appConfig)

	covidStatsStore := store.NewCovidStatsStore(lgr, dbConn)
	covidStatsService := service.NewCovidStatsService(lgr, covidStatsStore)
	covidStatsHandler := handler.NewCovidStatHandler(lgr, covidStatsService)
	api.InitCovidStatRoutes(router, covidStatsHandler)

	server.StartAndGracefullShutdown(lgr, router, appConfig.SERVER)
}
