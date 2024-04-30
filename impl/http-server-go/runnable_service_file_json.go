package main

import (
	"log/slog"
	"openapi"
	"rebootx-on-prem/http-server-go/config"
	"rebootx-on-prem/http-server-go/utils"
)

type RunnableServiceFileJson struct {
	config *config.Config
	logger *slog.Logger
}

func (service RunnableServiceFileJson) list(params *openapi.ListRunnablesQueryParams) (*openapi.ListResRunnable, *utils.ServiceError) {
	config := service.config

	items, err := utils.LoadItemsFromJson[openapi.Runnable](config.RunnableServiceFileJsonFilePath)
	if err != nil {
		return nil, err
	}

	total := int32(len(items))

	res := openapi.NewListResRunnable(items, total)

	return res, nil
}

func (service RunnableServiceFileJson) reboot(id string) (*openapi.RunnableOperationRes, *utils.ServiceError) {
	config := service.config
	logger := service.logger

	item, err := utils.LoadItemfromJson[openapi.Runnable](config.RunnableServiceFileJsonFilePath, func(r openapi.Runnable) bool { return r.Id == id })
	if err != nil {
		return nil, err
	}

	// In a real world, we would probably SSH into the instance to perform the command
	logger.Info("Faking reboot", "id", item.Id, "name", item.Name)

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}

func (service RunnableServiceFileJson) stop(id string) (*openapi.RunnableOperationRes, *utils.ServiceError) {
	config := service.config
	logger := service.logger

	item, err := utils.LoadItemfromJson[openapi.Runnable](config.RunnableServiceFileJsonFilePath, func(r openapi.Runnable) bool { return r.Id == id })
	if err != nil {
		return nil, err
	}

	// In a real world, we would probably SSH into the instance to perform the command
	logger.Info("Faking stop", "id", item.Id, "name", item.Name)

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}
