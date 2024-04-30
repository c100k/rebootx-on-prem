package main

import (
	"openapi"
	"rebootx-on-prem/http-server-go/config"
	"rebootx-on-prem/http-server-go/utils"
)

type DashboardServiceFileJson struct {
	config *config.Config
}

func (service DashboardServiceFileJson) list(params *openapi.ListDashboardsQueryParams) (*openapi.ListResDashboard, *utils.ServiceError) {
	config := service.config

	items, err := utils.LoadItemsFromJson[openapi.Dashboard](config.DashboardServiceFileJsonFilePath)
	if err != nil {
		return nil, err
	}

	total := int32(len(items))

	res := openapi.NewListResDashboard(items, total)

	return res, nil
}
