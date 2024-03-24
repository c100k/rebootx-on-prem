package main

import (
	"math"
	"strconv"
)

func parseInt(raw *string) *int32 {
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

func ptr[T float64 | string](v T) *T {
	return &v
}

func roundToCloser(v float64) float64 {
	return math.Round(v*100) / 100
}
