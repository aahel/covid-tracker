package api

import (
	"net/http"

	"github.com/aahel/covid-tracker/errors"
	"github.com/aahel/covid-tracker/respond"
)

// Handler custom api handler help us to handle all the errors in one place
type Handler func(w http.ResponseWriter, r *http.Request) *errors.AppError

func (f Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := f(w, r)

	if err != nil {
		respond.Fail(w, err)
	}
}
