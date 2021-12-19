package service

import (
	"fmt"

	"github.com/aahel/covid-tracker/config"
	"github.com/aahel/covid-tracker/entity"
	"github.com/aahel/covid-tracker/errors"
	"github.com/aahel/covid-tracker/store"
	"github.com/aahel/covid-tracker/types"
	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type CovidStatsServiceImpl struct {
	l     *zap.SugaredLogger
	store store.CovidStatsStore
}

func NewCovidStatsService(l *zap.SugaredLogger, db store.CovidStatsStore) *CovidStatsServiceImpl {
	return &CovidStatsServiceImpl{l, db}
}

func (csSvc *CovidStatsServiceImpl) SaveCovidStats() ([]*entity.CovidStat, *errors.AppError) {
	covidData, err := FetchCovidData()
	if err != nil {
		return nil, err
	}

	covidStats := []*entity.CovidStat{}
	for state, stat := range covidData {
		covidStats = append(covidStats, &entity.CovidStat{
			ID:            primitive.NewObjectID(),
			StateCode:     state,
			StateCount:    stat.Total.Confirmed,
			LastUpdatedAt: stat.Meta.LastUpdated,
		})
	}
	covidStats, errx := csSvc.store.SaveCovidInfo(covidStats)
	if errx != nil {
		return nil, errx
	}
	return covidStats, nil
}

func FetchCovidData() (map[string]types.CovidStat, *errors.AppError) {
	cvoidDataResp := map[string]types.CovidStat{}
	client := resty.New()
	resp, err := client.R().
		SetResult(&cvoidDataResp).
		Get(config.CovidApiUrl)
	if err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}
	if resp.StatusCode() <= 299 {
		return cvoidDataResp, nil
	}
	return nil, errors.InternalServerStd()
}

func (csSvc *CovidStatsServiceImpl) GetCovidStats(lat, long string) (*types.CovidStatResp, *errors.AppError) {
	stateCode, err := FetchStateCode(lat, long)
	if err != nil {
		return nil, err
	}
	covidStat, indiaCases, errx := csSvc.store.GetCovidInfo(stateCode)
	if errx != nil {
		return nil, errx
	}
	covidStatResp := &types.CovidStatResp{
		StateCode:     covidStat.StateCode,
		StateCount:    covidStat.StateCount,
		LastUpdatedAt: covidStat.LastUpdatedAt,
		IndiaCount:    indiaCases,
	}
	return covidStatResp, nil
}

func FetchStateCode(lat, long string) (string, *errors.AppError) {
	geoCodingData := &types.ReverseGeoCodingResp{}
	client := resty.New()
	resp, err := client.SetDebug(client.Debug).R().
		SetQueryParam("access_key", config.GetAppConfig().APIKEY.GeoCoding).
		SetQueryParam("query", fmt.Sprintf("%s,%s", lat, long)).
		SetQueryParam("limit", "1").
		SetResult(geoCodingData).
		Get(config.ReverseGeoCodingAPiUrl)
	if err != nil {
		return "", errors.InternalServerStd().AddDebug(err)
	}
	if resp.StatusCode() == 200 {
		return geoCodingData.Data[0].RegionCode, nil
	}
	if resp.StatusCode() <= 499 {
		return "", errors.NotFound("location not found")
	}
	return "", errors.InternalServerStd()
}
