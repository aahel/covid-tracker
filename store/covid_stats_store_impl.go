package store

import (
	"context"

	"github.com/aahel/covid-tracker/entity"
	"github.com/aahel/covid-tracker/errors"
	"github.com/aahel/covid-tracker/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type CovidStatsStoreImpl struct {
	l  *zap.SugaredLogger
	db *mongo.Database
}

func NewCovidStatsStore(l *zap.SugaredLogger, db *mongo.Database) *CovidStatsStoreImpl {
	return &CovidStatsStoreImpl{l, db}
}

func (csStr *CovidStatsStoreImpl) SaveCovidInfo(covidStats []*entity.CovidStat) ([]*entity.CovidStat, *errors.AppError) {
	var covidDocs []interface{}
	for _, cs := range covidStats {
		covidDocs = append(covidDocs, cs)
	}
	_, err := csStr.db.Collection("covid_stats").InsertMany(context.TODO(), covidDocs)
	if err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return covidStats, nil
}

func (csStr *CovidStatsStoreImpl) GetCovidInfo(stateCode string) (*entity.CovidStat, int64, *errors.AppError) {
	covidStat := &entity.CovidStat{}
	total := []*types.TotalCases{}
	filter := map[string]interface{}{"state_code": stateCode}
	options := options.FindOne()
	options.SetSort(bson.M{"_id": -1})
	if err := csStr.db.Collection("covid_stats").FindOne(context.TODO(), filter, options).Decode(covidStat); err != nil {
		if errors.IsMongoNoDocErr(err) {
			return nil, 0, errors.NotFound("covid data not found")
		}
		return nil, 0, errors.InternalServerStd().AddDebug(err)
	}
	pipeline := []map[string]interface{}{
		{
			"$group": map[string]interface{}{
				"_id":         "",
				"india_count": map[string]interface{}{"$sum": "$state_count"},
			},
		},
		{
			"$project": map[string]interface{}{
				"_id":         0,
				"total_cases": "$india_count",
			},
		},
	}
	ctx := context.Background()
	i, err := csStr.db.Collection("covid_stats").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, 0, errors.InternalServerStd().AddDebug(err)
	}
	defer i.Close(ctx)
	if err := i.All(ctx, &total); err != nil {
		return nil, 0, errors.InternalServerStd().AddDebug(err)
	}
	return covidStat, total[0].TotalCases, nil
}
