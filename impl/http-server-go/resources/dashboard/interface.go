package dashboard

import (
	"fmt"
	"openapi"
	"rebootx-on-prem/http-server-go/config"
	"rebootx-on-prem/http-server-go/resources/dashboard/file_json"
	"rebootx-on-prem/http-server-go/utils"
)

type Service interface {
	List(params *openapi.ListDashboardsQueryParams) (*openapi.ListResDashboard, *utils.ServiceError)
}

func LoadService(config *config.Config) *Service {
	var service Service
	switch config.DashboardServiceImpl {
	case "fileJson":
		service = file_json.NewService(config)
	default:
		panic(fmt.Sprintf("Invalid DashboardServiceImpl : %s", config.DashboardServiceImpl))
	}

	return &service
}
