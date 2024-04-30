package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
)

func LoadItemsFromJson[T any](filePath *string) ([]T, *ServiceError) {
	file, err := os.Open(*filePath)
	if err != nil {
		return nil, &ServiceError{HttpStatus: 500, Message: err.Error()}
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, &ServiceError{HttpStatus: 500, Message: err.Error()}
	}

	var items []T
	json.Unmarshal(content, &items)
	if items == nil {
		return nil, &ServiceError{HttpStatus: 500, Message: fmt.Sprintf("Fix the file %s to respect the expected schema", *filePath)}
	}

	return items, nil
}

func LoadItemfromJson[T any](filePath *string, predicate func(T) bool) (*T, *ServiceError) {
	items, err := LoadItemsFromJson[T](filePath)
	if err != nil {
		return nil, err
	}

	// If this happens to be called lots of times and we know the file is not changed after starting the server,
	// this can be optimized by creating a Map to search faster
	idx := slices.IndexFunc(items, predicate)
	if idx == -1 {
		return nil, &ServiceError{HttpStatus: 404, Message: Err404}
	}

	return &items[idx], nil
}

func ParseInt(raw *string) *int32 {
	if raw == nil {
		return nil
	}

	v, err := strconv.ParseInt(*raw, 0, 32)
	if err != nil {
		return nil
	}

	asInt32 := int32(v)

	return &asInt32
}

func Ptr[T float64 | string](v T) *T {
	return &v
}

func RoundToCloser(v float64) float64 {
	return math.Round(v*100) / 100
}
