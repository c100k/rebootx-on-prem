package main

import (
	"fmt"
	"openapi"
)

type DashboardService interface {
	list(params *openapi.ListDashboardsQueryParams) (*openapi.ListResDashboard, *ServiceError)
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
