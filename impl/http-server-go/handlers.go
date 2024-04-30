package main

import (
	"encoding/json"
	"net/http"
	resources_dashboard "rebootx-on-prem/http-server-go/resources/dashboard"
	resources_runnable "rebootx-on-prem/http-server-go/resources/runnable"
	"rebootx-on-prem/http-server-go/utils"

	"github.com/gorilla/mux"

	"openapi"
)

func getDashboardsHandler(service resources_dashboard.Service) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)

		params := openapi.NewListDashboardsQueryParams()
		query := r.URL.Query()
		limitAsString := query.Get("limit")
		limit := utils.ParseInt(&limitAsString)
		if limit != nil {
			params.SetLimit(*limit)
		}
		offsetAsString := query.Get("offset")
		offset := utils.ParseInt(&offsetAsString)
		if offset != nil {
			params.SetOffset(*offset)
		}
		params.SetQ(query.Get("q"))

		res, err := service.List(params)
		if err != nil {
			w.WriteHeader(err.HTTPStatus)
			encoder.Encode(openapi.NewErrorRes(err.Error()))
			return
		}

		encoder.Encode(res)
	})
}

func getRunnablesHandler(service resources_runnable.Service) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)

		params := openapi.NewListRunnablesQueryParams()
		query := r.URL.Query()
		limitAsString := query.Get("limit")
		limit := utils.ParseInt(&limitAsString)
		if limit != nil {
			params.SetLimit(*limit)
		}
		offsetAsString := query.Get("offset")
		offset := utils.ParseInt(&offsetAsString)
		if offset != nil {
			params.SetOffset(*offset)
		}
		params.SetQ(query.Get("q"))

		res, err := service.List(params)
		if err != nil {
			w.WriteHeader(err.HTTPStatus)
			encoder.Encode(openapi.NewErrorRes(err.Error()))
			return
		}

		encoder.Encode(res)
	})
}

func postRunnableRebootHandler(service resources_runnable.Service) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)

		vars := mux.Vars(r)
		id := vars["id"]

		res, err := service.Reboot(id)
		if err != nil {
			w.WriteHeader(err.HTTPStatus)
			encoder.Encode(openapi.NewErrorRes(err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)
		encoder.Encode(res)
	})
}

func postRunnableStopHandler(service resources_runnable.Service) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)

		vars := mux.Vars(r)
		id := vars["id"]

		res, err := service.Stop(id)
		if err != nil {
			w.WriteHeader(err.HTTPStatus)
			encoder.Encode(openapi.NewErrorRes(err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)
		encoder.Encode(res)
	})
}
