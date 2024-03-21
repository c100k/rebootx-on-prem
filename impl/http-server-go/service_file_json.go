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

func (service ServiceFileJson) list(params *openapi.ListRunnablesQueryParams) (*ServiceError, *openapi.ListResRunnable) {
	config := service.config

	err, items := findItems(config)
	if err != nil {
		return err, nil
	}

	total := int32(len(items))

	res := openapi.NewListResRunnable(items, total)

	return nil, res
}

func (service ServiceFileJson) reboot(id string) (*ServiceError, *openapi.RunnableOperationRes) {
	config := service.config
	logger := service.logger

	err, item := findItem(config, id)
	if err != nil {
		return err, nil
	}

	// In a real world, we would probably SSH into the instance to perform the command
	logger.Info("Faking reboot", "id", item.Id, "name", item.Name)

	return nil, openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil))
}

func (service ServiceFileJson) stop(id string) (*ServiceError, *openapi.RunnableOperationRes) {
	config := service.config
	logger := service.logger

	err, item := findItem(config, id)
	if err != nil {
		return err, nil
	}

	// In a real world, we would probably SSH into the instance to perform the command
	logger.Info("Faking stop", "id", item.Id, "name", item.Name)

	return nil, openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil))
}

func findItems(config *Config) (*ServiceError, []openapi.Runnable) {
	file, err := os.Open(*config.serviceFileJsonFilePath)
	if err != nil {
		return &ServiceError{HttpStatus: 500, Message: err.Error()}, nil
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return &ServiceError{HttpStatus: 500, Message: err.Error()}, nil
	}

	var items []openapi.Runnable
	json.Unmarshal(content, &items)
	if items == nil {
		return &ServiceError{HttpStatus: 500, Message: "Fix your JSON file to respect the schema"}, nil
	}

	return nil, items
}

func findItem(config *Config, id string) (*ServiceError, *openapi.Runnable) {
	err, items := findItems(config)
	if err != nil {
		return err, nil
	}

	// If this happens to be called lots of times and we know the file is not changed after starting the server,
	// this can be optimized by creating a Map to search faster
	idx := slices.IndexFunc(items, func(i openapi.Runnable) bool { return i.Id == id })
	if idx == -1 {
		return &ServiceError{HttpStatus: 404, Message: Err404Runnable}, nil
	}

	return nil, &items[idx]
}
