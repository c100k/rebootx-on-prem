package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"openapi"
)

func getRunnablesHandler(service RunnableService) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)

		params := openapi.NewListRunnablesQueryParams()
		query := r.URL.Query()
		limitAsString := query.Get("limit")
		limit := parseInt(&limitAsString)
		if limit != nil {
			params.SetLimit(*limit)
		}
		params.SetQ(query.Get("q"))
		offsetAsString := query.Get("offset")
		offset := parseInt(&offsetAsString)
		if offset != nil {
			params.SetOffset(*offset)
		}

		res, err := service.list(params)
		if err != nil {
			w.WriteHeader(err.HttpStatus)
			encoder.Encode(openapi.NewErrorRes(err.Error()))
			return
		}

		encoder.Encode(res)
	})
}

func postRunnableRebootHandler(service RunnableService) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)

		vars := mux.Vars(r)
		id := vars["id"]

		res, err := service.reboot(id)
		if err != nil {
			w.WriteHeader(err.HttpStatus)
			encoder.Encode(openapi.NewErrorRes(err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)
		encoder.Encode(res)
	})
}

func postRunnableStopHandler(service RunnableService) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)

		vars := mux.Vars(r)
		id := vars["id"]

		res, err := service.stop(id)
		if err != nil {
			w.WriteHeader(err.HttpStatus)
			encoder.Encode(openapi.NewErrorRes(err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)
		encoder.Encode(res)
	})
}
