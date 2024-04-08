package main

import (
	"log/slog"
	"openapi"
)

type RunnableServiceFileJson struct {
	config *Config
	logger *slog.Logger
}

func (service RunnableServiceFileJson) list(params *openapi.ListRunnablesQueryParams) (*openapi.ListResRunnable, *ServiceError) {
	config := service.config

	items, err := loadItemsFromJson[openapi.Runnable](config.runnableServiceFileJsonFilePath)
	if err != nil {
		return nil, err
	}

	total := int32(len(items))

	res := openapi.NewListResRunnable(items, total)

	return res, nil
}

func (service RunnableServiceFileJson) reboot(id string) (*openapi.RunnableOperationRes, *ServiceError) {
	config := service.config
	logger := service.logger

	item, err := loadItemfromJson[openapi.Runnable](config.runnableServiceFileJsonFilePath, func(r openapi.Runnable) bool { return r.Id == id })
	if err != nil {
		return nil, err
	}

	// In a real world, we would probably SSH into the instance to perform the command
	logger.Info("Faking reboot", "id", item.Id, "name", item.Name)

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}

func (service RunnableServiceFileJson) stop(id string) (*openapi.RunnableOperationRes, *ServiceError) {
	config := service.config
	logger := service.logger

	item, err := loadItemfromJson[openapi.Runnable](config.runnableServiceFileJsonFilePath, func(r openapi.Runnable) bool { return r.Id == id })
	if err != nil {
		return nil, err
	}

	// In a real world, we would probably SSH into the instance to perform the command
	logger.Info("Faking stop", "id", item.Id, "name", item.Name)

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}
