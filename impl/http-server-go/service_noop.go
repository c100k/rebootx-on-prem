package main

import (
	"log/slog"
	"openapi"
)

type ServiceNoop struct {
	logger *slog.Logger
}

func (service ServiceNoop) list(params *openapi.ListRunnablesQueryParams) (*openapi.ListResRunnable, *ServiceError) {
	service.logger.Warn("Noop")

	items := []openapi.Runnable{}
	total := int32(len(items))

	res := openapi.NewListResRunnable(items, total)

	return res, nil
}

func (service ServiceNoop) reboot(id string) (*openapi.RunnableOperationRes, *ServiceError) {
	service.logger.Warn("Noop")

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}

func (service ServiceNoop) stop(id string) (*openapi.RunnableOperationRes, *ServiceError) {
	service.logger.Warn("Noop")

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}
