package main

import (
	"openapi"
)

type DashboardServiceFileJson struct {
	config *Config
}

func (service DashboardServiceFileJson) getMetric(dashboardId string, metricId string) (*openapi.DashboardMetric, *ServiceError) {
	return nil, nil
}

func (service DashboardServiceFileJson) list(params *openapi.ListDashboardsQueryParams) (*openapi.ListResDashboard, *ServiceError) {
	config := service.config

	items, err := loadItemsFromJson[openapi.Dashboard](config.dashboardServiceFileJsonFilePath)
	if err != nil {
		return nil, err
	}

	total := int32(len(items))

	res := openapi.NewListResDashboard(items, total)

	return res, nil
}

func (service DashboardServiceFileJson) listMetrics(dashboardId string) (*openapi.ListResDashboardMetric, *ServiceError) {
	return nil, nil
}
