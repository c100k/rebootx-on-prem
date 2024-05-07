package file_json

import (
	"log/slog"
	"openapi"
	"os"
	"rebootx-on-prem/http-server-go/config"
	"rebootx-on-prem/http-server-go/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	// Given
	config := config.New()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	config.RunnableServiceFileJsonFilePath = utils.Ptr("../../../../../data/servers.example.json")
	service := NewService(config, logger)
	params := openapi.NewListRunnablesQueryParamsWithDefaults()

	// When
	res, err := service.List(params)
	items := res.Items

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
