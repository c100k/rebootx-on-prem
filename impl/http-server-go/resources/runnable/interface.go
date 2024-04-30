package resources_runnable

import (
	"fmt"
	"log/slog"
	"openapi"
	"rebootx-on-prem/http-server-go/config"
	"rebootx-on-prem/http-server-go/resources/runnable/file_json"
	"rebootx-on-prem/http-server-go/resources/runnable/self"
	"rebootx-on-prem/http-server-go/utils"
)

type Service interface {
	List(params *openapi.ListRunnablesQueryParams) (*openapi.ListResRunnable, *utils.ServiceError)
	Reboot(id string) (*openapi.RunnableOperationRes, *utils.ServiceError)
	Stop(id string) (*openapi.RunnableOperationRes, *utils.ServiceError)
}

func LoadService(config *config.Config, logger *slog.Logger) *Service {
	var service Service
	switch config.RunnableServiceImpl {
	case "fileJson":
		service = file_json.NewService(config, logger)
	case "self":
		service = self.NewService(config, logger)
	default:
		panic(fmt.Sprintf("Invalid RunnableServiceImpl : %s", config.RunnableServiceImpl))
	}

	return &service
}
