package api

import (
	"net/http"

	v1 "github.com/aahel/covid-tracker/api/v1"
	"github.com/go-chi/chi"
)

func InitCovidStatRoutes(r chi.Router, cov *v1.CovidStatHandler) {
	r.Route("/v1/covidstats", func(r chi.Router) {
		r.Method(http.MethodPost, "/", Handler(cov.SaveCovidStats))
		r.Method(http.MethodGet, "/", Handler(cov.GetCovidStats))
	})
}
