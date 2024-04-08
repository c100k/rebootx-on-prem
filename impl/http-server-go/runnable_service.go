package main

import (
	"fmt"
	"log/slog"
	"openapi"
)

type RunnableServiceOperationType int32

const (
	REBOOT RunnableServiceOperationType = 0
	STOP   RunnableServiceOperationType = 1
)

type RunnableService interface {
	list(params *openapi.ListRunnablesQueryParams) (*openapi.ListResRunnable, *ServiceError)
	reboot(id string) (*openapi.RunnableOperationRes, *ServiceError)
	stop(id string) (*openapi.RunnableOperationRes, *ServiceError)
}

func loadRunnableService(config *Config, logger *slog.Logger) *RunnableService {
	var service RunnableService
	switch config.runnableServiceImpl {
	case "fileJson":
		service = RunnableServiceFileJson{config: config, logger: logger}
	case "noop":
		service = RunnableServiceNoop{logger: logger}
	case "self":
		service = RunnableServiceSelf{config: config, logger: logger}
	default:
		panic(fmt.Sprintf("Invalid runnableServiceImpl : %s", config.runnableServiceImpl))
	}

	return &service
}
