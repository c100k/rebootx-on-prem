package main

import "openapi"

type ServiceOperationType int32

const (
	REBOOT ServiceOperationType = 0
	STOP   ServiceOperationType = 1
)

type Service interface {
	list(params *openapi.ListRunnablesQueryParams) (*openapi.ListResRunnable, *ServiceError)
	reboot(id string) (*openapi.RunnableOperationRes, *ServiceError)
	stop(id string) (*openapi.RunnableOperationRes, *ServiceError)
}
