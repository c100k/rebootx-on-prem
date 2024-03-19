package main

import "openapi"

type Service interface {
	list(params *openapi.ListRunnablesQueryParams) (*ServiceError, *openapi.ListResRunnable)
	reboot(id string) (*ServiceError, *openapi.RunnableOperationRes)
	stop(id string) (*ServiceError, *openapi.RunnableOperationRes)
}
