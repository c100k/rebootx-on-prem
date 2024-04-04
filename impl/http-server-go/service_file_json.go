package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"openapi"
	"os"
	"slices"
)

type ServiceFileJson struct {
	config *Config
	logger *slog.Logger
}

func (service ServiceFileJson) list(params *openapi.ListRunnablesQueryParams) (*openapi.ListResRunnable, *ServiceError) {
	config := service.config

	items, err := findItems(config.runnableServiceFileJsonFilePath)
	if err != nil {
		return nil, err
	}

	total := int32(len(items))

	res := openapi.NewListResRunnable(items, total)

	return res, nil
}

func (service ServiceFileJson) reboot(id string) (*openapi.RunnableOperationRes, *ServiceError) {
	config := service.config
	logger := service.logger

	item, err := findItem(config.runnableServiceFileJsonFilePath, id)
	if err != nil {
		return nil, err
	}

	// In a real world, we would probably SSH into the instance to perform the command
	logger.Info("Faking reboot", "id", item.Id, "name", item.Name)

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}

func (service ServiceFileJson) stop(id string) (*openapi.RunnableOperationRes, *ServiceError) {
	config := service.config
	logger := service.logger

	item, err := findItem(config.runnableServiceFileJsonFilePath, id)
	if err != nil {
		return nil, err
	}

	// In a real world, we would probably SSH into the instance to perform the command
	logger.Info("Faking stop", "id", item.Id, "name", item.Name)

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}

func findItems(filePath *string) ([]openapi.Runnable, *ServiceError) {
	file, err := os.Open(*filePath)
	if err != nil {
		return nil, &ServiceError{HttpStatus: 500, Message: err.Error()}
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, &ServiceError{HttpStatus: 500, Message: err.Error()}
	}

	var items []openapi.Runnable
	json.Unmarshal(content, &items)
	if items == nil {
		return nil, &ServiceError{HttpStatus: 500, Message: "Fix your JSON file to respect the schema"}
	}

	return items, nil
}

func findItem(filePath *string, id string) (*openapi.Runnable, *ServiceError) {
	items, err := findItems(filePath)
	if err != nil {
		return nil, err
	}

	// If this happens to be called lots of times and we know the file is not changed after starting the server,
	// this can be optimized by creating a Map to search faster
	idx := slices.IndexFunc(items, func(i openapi.Runnable) bool { return i.Id == id })
	if idx == -1 {
		return nil, &ServiceError{HttpStatus: 404, Message: Err404Runnable}
	}

	return &items[idx], nil
}
