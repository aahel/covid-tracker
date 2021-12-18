package store

import (
	"github.com/aahel/covid-tracker/entity"
	"github.com/aahel/covid-tracker/errors"
)

type CovidStatsStore interface {
	SaveCovidInfo(covidStats []*entity.CovidStat) ([]*entity.CovidStat, *errors.AppError)
	GetCovidInfo(stateCode string) (*entity.CovidStat, int64, *errors.AppError)
}
