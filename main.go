package main

import (
	"github.com/aahel/covid-tracker/api"
	handler "github.com/aahel/covid-tracker/api/v1"
	"github.com/aahel/covid-tracker/config"
	"github.com/aahel/covid-tracker/server"
	"github.com/aahel/covid-tracker/service"
	"github.com/aahel/covid-tracker/store"
	"github.com/go-openapi/runtime/middleware"
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

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	dh := middleware.Redoc(opts, nil)
	api.InitDocRoutes(router, dh)

	server.StartAndGracefullShutdown(lgr, router, appConfig.SERVER)
}
