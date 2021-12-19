package v1

import (
	"github.com/aahel/covid-tracker/entity"
	"github.com/aahel/covid-tracker/types"
)

type getpayload struct {
	// Data holds the response data
	Data types.CovidStatResp `json:"data"`
}

type InterServerErr struct {
	// example: 500
	Status int `json:"status"`
	// Error message
	// example: Something went wrong
	Message string `json:"error_message,omitempty"`
	// For future use ,not implemented ye
	Error interface{} `json:"error,omitempty"`
}

type CovidStatNotFoundErr struct {
	// example: 404
	Status int `json:"status"`
	// Error message
	// example: covid data not found
	Message string `json:"error_message,omitempty"`
}

type ErrIntServerWrapper struct {
	Error InterServerErr `json:"error"`
}

type ErrCovidStatNotFoundWrapper struct {
	Error CovidStatNotFoundErr `json:"error"`
}

// A list of covid stats
// swagger:response CovidStatSuccessResponse
type CovidStatSuccessResponse struct {
	// covid stat
	// in: body
	Body getpayload
}

// swagger:response CovidStatIntserverResponse
type CovidStatIntserverResponse struct {
	// in: body
	Body ErrIntServerWrapper
}

// swagger:response CovidStatNotFoundResponse
type CovidStatNotFoundResponse struct {
	// in: body
	Body ErrCovidStatNotFoundWrapper
}

// swagger:parameters getCovidStat
type CovidStatQueryParams struct {
	// in: query
	// required: true
	Latitude string `json:"latitude"`
	// required: true
	Longitude string `json:"longitude"`
}

type createSuceessPayload struct {
	Data []*entity.CovidStat `json:"data"`
}

// swagger:response CovidStatCreateSuccessResp
type SubjectCreateSuccessResp struct {
	//in: body
	Body createSuceessPayload
}
