package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"openapi"
	"os"
)

type ServiceFileJson struct {
	config *Config
	logger *slog.Logger
}

func (service ServiceFileJson) list(params *openapi.ListRunnablesQueryParams) (*ServiceError, *openapi.ListResRunnable) {
	config := service.config

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

	total := int32(len(items))

	res := openapi.NewListResRunnable(items, total)

	return nil, res
}

func (service ServiceFileJson) reboot(id string) (*ServiceError, *openapi.RunnableOperationRes) {
	return nil, openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil))
}

func (service ServiceFileJson) stop(id string) (*ServiceError, *openapi.RunnableOperationRes) {
	return nil, openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil))
}
