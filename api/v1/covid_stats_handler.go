package v1

import (
	"net/http"
	"strings"

	"github.com/aahel/covid-tracker/errors"
	"github.com/aahel/covid-tracker/respond"
	"github.com/aahel/covid-tracker/service"
	"github.com/aahel/covid-tracker/types"
	"go.uber.org/zap"
)

type CovidStatHandler struct {
	lgr *zap.SugaredLogger
	svc service.CovidStatsService
}

func NewCovidStatHandler(lgr *zap.SugaredLogger, svc service.CovidStatsService) *CovidStatHandler {
	return &CovidStatHandler{lgr, svc}
}

// swagger:route POST /v1/covidstats covidStat createCovidStats
// Create new covidStats
// responses:
//	201: CovidStatCreateSuccessResp
//	500: CovidStatIntserverResponse
func (cov *CovidStatHandler) SaveCovidStats(rw http.ResponseWriter, r *http.Request) *errors.AppError {
	covidStats, err := cov.svc.SaveCovidStats()
	if err != nil {
		return err
	}
	return respond.Created(rw, covidStats, nil)
}

// swagger:route GET /v1/covidstats covidStat getCovidStat
// Return a covidStat from the database
// responses:
//	200: CovidStatSuccessResponse
//	500: CovidStatIntserverResponse
//  404: CovidStatNotFoundResponse
func (cov *CovidStatHandler) GetCovidStats(rw http.ResponseWriter, r *http.Request) *errors.AppError {
	lat := r.URL.Query().Get(types.Latitude)
	long := r.URL.Query().Get(types.Longitude)
	if strings.TrimSpace(lat) == "" {
		return errors.BadRequest(errors.IsRequiredErr(types.Latitude).Error())
	}
	if strings.TrimSpace(long) == "" {
		return errors.BadRequest(errors.IsRequiredErr(types.Longitude).Error())
	}
	covidStats, err := cov.svc.GetCovidStats(lat, long)
	if err != nil {
		return err
	}

	return respond.OK(rw, covidStats, nil)
}
