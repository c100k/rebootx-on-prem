package main

import (
	"openapi"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDashboardsFromJson(t *testing.T) {
	// Given
	filePath := "../../data/dashboards.example.json"

	// When
	items, err := loadItemsFromJson[openapi.Dashboard](&filePath)

	// Then
	assert.Nil(t, err)
	assert.Len(t, items, 2)
	item := items[0]
	assert.Equal(t, "123", item.Id)
	assert.Equal(t, "Infra", item.Name)
}
