package main

import (
	"fmt"
	"log/slog"
	"openapi"
	"rebootx-on-prem/http-server-go/config"
	"rebootx-on-prem/http-server-go/utils"
)

type RunnableServiceOperationType int32

const (
	REBOOT RunnableServiceOperationType = 0
	STOP   RunnableServiceOperationType = 1
)

type RunnableService interface {
	list(params *openapi.ListRunnablesQueryParams) (*openapi.ListResRunnable, *utils.ServiceError)
	reboot(id string) (*openapi.RunnableOperationRes, *utils.ServiceError)
	stop(id string) (*openapi.RunnableOperationRes, *utils.ServiceError)
}

func loadRunnableService(config *config.Config, logger *slog.Logger) *RunnableService {
	var service RunnableService
	switch config.RunnableServiceImpl {
	case "fileJson":
		service = RunnableServiceFileJson{config: config, logger: logger}
	case "self":
		service = RunnableServiceSelf{config: config, logger: logger}
	default:
		panic(fmt.Sprintf("Invalid runnableServiceImpl : %s", config.RunnableServiceImpl))
	}

	return &service
}
