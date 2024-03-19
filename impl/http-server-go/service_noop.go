package main

import (
	"log/slog"
	"openapi"
)

type ServiceNoop struct {
	logger *slog.Logger
}

func (service ServiceNoop) list(params *openapi.ListRunnablesQueryParams) (*ServiceError, *openapi.ListResRunnable) {
	service.logger.Warn("Noop")

	items := []openapi.Runnable{}
	total := int32(len(items))

	res := openapi.NewListResRunnable(items, total)

	return nil, res
}

func (service ServiceNoop) reboot(id string) (*ServiceError, *openapi.RunnableOperationRes) {
	service.logger.Warn("Noop")

	return nil, openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil))
}

func (service ServiceNoop) stop(id string) (*ServiceError, *openapi.RunnableOperationRes) {
	service.logger.Warn("Noop")

	return nil, openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil))
}
