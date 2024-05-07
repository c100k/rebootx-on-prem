package self

import (
	"log/slog"
	"openapi"
	"os"
	"rebootx-on-prem/http-server-go/config"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	// Given
	config := config.New()
	config.RunnableServiceImpl = "self"
	config.RunnableServiceSelfID = "self"
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	service := NewService(config, logger)
	params := openapi.NewListRunnablesQueryParamsWithDefaults()

	// When
	res, err := service.List(params)
	items := res.Items

	// Then
	expectedName, _ := os.Hostname()
	assert.Nil(t, err)
	assert.Len(t, items, 1)
	item := items[0]
	assert.Nil(t, item.Flavor.Get())
	assert.Nil(t, item.Fqdn.Get())
	assert.Equal(t, "self", item.Id)
	assert.Len(t, item.Metrics, 3)
	assert.Equal(t, expectedName, item.Name)
	assert.Nil(t, item.Scopes.Geo.Get())
	assert.Nil(t, item.Scopes.Logical.Get())
	assert.Equal(t, int32(0), item.Ssh.Get().Port)
	assert.Nil(t, item.Stack.Get())
	assert.Equal(t, openapi.ON, item.Status)
}

func TestBuildCPUMetric(t *testing.T) {
	// Given
	used := uint64(23)
	total := uint64(1201)

	// When
	metric := buildCPUMetric(used, total)
	value := *metric.Value.Get()

	// Then
	assert.Equal(t, 1.92, value)
}

func TestBuildMemoryMetric(t *testing.T) {
	// Given
	used := uint64(14333321216)
	total := uint64(17176596480)

	// When
	metric := buildMemoryMetric(used, total)
	ratio := *metric.Ratio.Get()
	thresholds := metric.Thresholds
	value := *metric.Value.Get()

	// Then
	assert.Equal(t, 0.83, ratio)
	assert.EqualValues(t, []float64{1288.24, 1460.01}, thresholds)
	assert.Equal(t, 1433.33, value)
}

func TestBuildUptimeMetric(t *testing.T) {
	// Given
	uptime, _ := time.ParseDuration("0h34m")

	// When
	metric := buildUptimeMetric(uptime)
	value := *metric.Value.Get()
	unit := *metric.Unit.Get()

	// Then
	assert.Equal(t, 34.0, value)
	assert.Equal(t, "min", unit)

	// Given
	uptime, _ = time.ParseDuration("2h34m")

	// When
	metric = buildUptimeMetric(uptime)
	value = *metric.Value.Get()
	unit = *metric.Unit.Get()

	// Then
	assert.Equal(t, 3.00, value)
	assert.Equal(t, "h", unit)
}
