package main

import (
	"fmt"
	"openapi"
)

type DashboardService interface {
	getMetric(dashboardId string, metricId string) (*openapi.DashboardMetric, *ServiceError)
	list(params *openapi.ListDashboardsQueryParams) (*openapi.ListResDashboard, *ServiceError)
	listMetrics(dashboardId string) (*openapi.ListResDashboardMetric, *ServiceError)
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
