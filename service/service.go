package service

import (
	"github.com/aahel/covid-tracker/entity"
	"github.com/aahel/covid-tracker/errors"
	"github.com/aahel/covid-tracker/types"
)

type CovidStatsService interface {
	SaveCovidStats() ([]*entity.CovidStat, *errors.AppError)
	GetCovidStats(lat, long string) (*types.CovidStatResp, *errors.AppError)
}
