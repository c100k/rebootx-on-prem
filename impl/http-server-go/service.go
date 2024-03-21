package main

import "openapi"

type Service interface {
	list(params *openapi.ListRunnablesQueryParams) (*openapi.ListResRunnable, *ServiceError)
	reboot(id string) (*openapi.RunnableOperationRes, *ServiceError)
	stop(id string) (*openapi.RunnableOperationRes, *ServiceError)
}
