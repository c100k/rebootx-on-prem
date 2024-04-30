package file_json

import (
	"openapi"
	"rebootx-on-prem/http-server-go/config"
	"rebootx-on-prem/http-server-go/utils"
)

type Service struct {
	config *config.Config
}

func NewService(config *config.Config) *Service {
	return &Service{config: config}
}

func (service Service) List(params *openapi.ListDashboardsQueryParams) (*openapi.ListResDashboard, *utils.ServiceError) {
	config := service.config

	items, err := utils.LoadItemsFromJson[openapi.Dashboard](config.DashboardServiceFileJsonFilePath)
	if err != nil {
		return nil, err
	}

	total := int32(len(items))

	res := openapi.NewListResDashboard(items, total)

	return res, nil
}
