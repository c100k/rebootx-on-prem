package main

import (
	"fmt"
	"openapi"
	"rebootx-on-prem/http-server-go/utils"
)

type DashboardService interface {
	list(params *openapi.ListDashboardsQueryParams) (*openapi.ListResDashboard, *utils.ServiceError)
}

func loadDashboardService(config *Config) *DashboardService {
	var service DashboardService
	switch config.dashboardServiceImpl {
	case "fileJson":
		service = DashboardServiceFileJson{config: config}
	default:
		panic(fmt.Sprintf("Invalid dashboardServiceImpl : %s", config.dashboardServiceImpl))
	}

	return &service
}
