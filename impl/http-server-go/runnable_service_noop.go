package main

import (
	"log/slog"
	"openapi"
)

type RunnableServiceNoop struct {
	logger *slog.Logger
}

func (service RunnableServiceNoop) list(params *openapi.ListRunnablesQueryParams) (*openapi.ListResRunnable, *ServiceError) {
	service.logger.Warn("Noop")

	items := []openapi.Runnable{}
	total := int32(len(items))

	res := openapi.NewListResRunnable(items, total)

	return res, nil
}

func (service RunnableServiceNoop) reboot(id string) (*openapi.RunnableOperationRes, *ServiceError) {
	service.logger.Warn("Noop")

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}

func (service RunnableServiceNoop) stop(id string) (*openapi.RunnableOperationRes, *ServiceError) {
	service.logger.Warn("Noop")

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}
