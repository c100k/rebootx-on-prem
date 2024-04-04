package main

import "openapi"

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
