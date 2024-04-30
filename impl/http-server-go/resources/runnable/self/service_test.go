package self

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
