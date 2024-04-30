package resources_dashboard

import (
	"fmt"
	"openapi"
	"rebootx-on-prem/http-server-go/config"
	"rebootx-on-prem/http-server-go/utils"
)

type DashboardService interface {
	List(params *openapi.ListDashboardsQueryParams) (*openapi.ListResDashboard, *utils.ServiceError)
}

func LoadDashboardService(config *config.Config) *DashboardService {
	var service DashboardService
	switch config.DashboardServiceImpl {
	case "fileJson":
		service = DashboardServiceFileJson{config: config}
	default:
		panic(fmt.Sprintf("Invalid dashboardServiceImpl : %s", config.DashboardServiceImpl))
	}

	return &service
}
