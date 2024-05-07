package file_json

import (
	"openapi"
	"rebootx-on-prem/http-server-go/config"
	"rebootx-on-prem/http-server-go/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	// Given
	config := config.New()
	config.DashboardServiceFileJsonFilePath = utils.Ptr("../../../../../data/dashboards.example.json")
	service := NewService(config)
	params := openapi.NewListDashboardsQueryParamsWithDefaults()

	// When
	res, err := service.List(params)
	items := res.Items

	// Then
	assert.Nil(t, err)
	assert.Len(t, items, 2)
	item0 := items[0]
	assert.Equal(t, "123", item0.Id)
	assert.Len(t, item0.Metrics, 0)
	assert.Equal(t, "Infra", item0.Name)
	item1 := items[1]
	assert.Equal(t, "456", item1.Id)
	assert.Len(t, item1.Metrics, 3)
	metric10 := item1.Metrics[0]
	assert.Equal(t, "123", metric10.Id)
	assert.Equal(t, "Clients #", *metric10.Label.Get())
	assert.Nil(t, metric10.Unit.Get())
	assert.Equal(t, 612.0, *metric10.Value.Get())
	assert.Equal(t, "Business", item1.Name)
}
