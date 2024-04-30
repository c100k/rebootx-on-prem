package resources_runnable

import (
	"openapi"
	"rebootx-on-prem/http-server-go/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadRunnablesFromJson(t *testing.T) {
	// Given
	filePath := "../../data/servers.example.json"

	// When
	items, err := utils.LoadItemsFromJson[openapi.Runnable](&filePath)

	// Then
	assert.Nil(t, err)
	assert.Len(t, items, 2)
	item := items[0]
	assert.Equal(t, "medium", *item.Flavor.Get())
	assert.Equal(t, "server01.mycompany.com", *item.Fqdn.Get())
	assert.Equal(t, "123", item.Id)
	assert.Len(t, item.Metrics, 2)
	assert.Equal(t, "server01", item.Name)
	assert.Equal(t, "Paris 01", item.Scopes.Geo.Get().Label)
	assert.Equal(t, "Project 1", item.Scopes.Logical.Get().Label)
	assert.Equal(t, int32(22), item.Ssh.Get().Port)
	assert.Equal(t, "nodejs", *item.Stack.Get())
	assert.Equal(t, openapi.OFF, item.Status)
}
