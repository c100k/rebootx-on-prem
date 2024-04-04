package main

import (
	"log/slog"
	"math"
	"openapi"
	"os"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/uptime"
)

type RunnableServiceSelf struct {
	config *Config
	logger *slog.Logger
}

const CPU_METRIC_LABEL = "CPU"
const CPU_METRIC_UNIT = "%"
const MEMORY_METRIC_LABEL = "RAM"
const MEMORY_METRIC_UNIT = "MB"
const MEMORY_METRIC_UNIT_AS_BYTES = 10e6
const THRESHOLD_WARNING = 0.75
const THRESHOLD_DANGER = 0.85
const UPTIME_METRIC_LABEL = "Uptime"

func (service RunnableServiceSelf) list(params *openapi.ListRunnablesQueryParams) (*openapi.ListResRunnable, *ServiceError) {
	config := service.config

	q := params.Q
	if len(*q) > 0 {
		err := checkThatRunnableExists(config, *q)
		if err != nil {
			return openapi.NewListResRunnable([]openapi.Runnable{}, 0), nil
		}
	}

	metrics := []openapi.RunnableMetric{}

	cpuValue, cpuTotal, err := getCPUStats()
	if err == nil {
		metrics = append(metrics, *buildCPUMetric(*cpuValue, *cpuTotal))
	}

	memory, err := memory.Get()
	if err == nil {
		metrics = append(metrics, *buildMemoryMetric(memory.Used, memory.Total))
	}

	uptime, err := uptime.Get()
	if err == nil {
		metrics = append(metrics, *buildUptimeMetric(uptime))
	}

	items := []openapi.Runnable{
		*openapi.NewRunnable(
			*getNullableFrom(config.runnableServiceSelfFlavor),
			*getNullableFrom(config.runnableServiceSelfFQDN),
			config.runnableServiceSelfId,
			*getNullableFrom(config.runnableServiceSelfIPv4),
			metrics,
			getNameFromHostname(config),
			*openapi.NewRunnableScopes(
				*getScope(config.runnableServiceSelfScopesGeoLabel, config.runnableServiceSelfScopesGeoValue),
				*getScope(config.runnableServiceSelfScopesLogicalLabel, config.runnableServiceSelfScopesLogicalValue),
			),
			*openapi.NewNullableRunnableSSH(
				openapi.NewRunnableSSH(
					*openapi.NewNullableString(&config.runnableServiceSelfSSHKeyname),
					config.runnableServiceSelfSSHPort,
					config.runnableServiceSelfSSHUsername,
				),
			),
			*getNullableFrom(config.runnableServiceSelfStack),
			openapi.ON,
		),
	}
	total := int32(len(items))

	res := openapi.NewListResRunnable(items, total)

	return res, nil
}

func (service RunnableServiceSelf) reboot(id string) (*openapi.RunnableOperationRes, *ServiceError) {
	config := service.config

	err := checkThatRunnableExists(config, id)
	if err != nil {
		return nil, err
	}

	errExec := performOpOnSelf(config, REBOOT)
	if errExec != nil {
		return nil, &ServiceError{HttpStatus: 500, Message: errExec.Error()}
	}

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}

func (service RunnableServiceSelf) stop(id string) (*openapi.RunnableOperationRes, *ServiceError) {
	config := service.config

	err := checkThatRunnableExists(config, id)
	if err != nil {
		return nil, err
	}

	errExec := performOpOnSelf(config, STOP)
	if errExec != nil {
		return nil, &ServiceError{HttpStatus: 500, Message: errExec.Error()}
	}

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}

func buildCPUMetric(used uint64, total uint64) *openapi.RunnableMetric {
	value := roundToCloser(float64(used) / float64(total) * 100)

	metric := openapi.NewRunnableMetric(
		*openapi.NewNullableString(ptr(CPU_METRIC_LABEL)),
		*openapi.NewNullableFloat64(nil),
		[]float64{THRESHOLD_WARNING, THRESHOLD_DANGER},
		*openapi.NewNullableString(ptr(CPU_METRIC_UNIT)),
		*openapi.NewNullableFloat64(&value),
	)

	return metric
}

func buildMemoryMetric(used uint64, total uint64) *openapi.RunnableMetric {
	valueInBytes := float64(used)

	ratio := roundToCloser(valueInBytes / float64(total))
	warning := roundToCloser(THRESHOLD_WARNING * float64(total) / MEMORY_METRIC_UNIT_AS_BYTES)
	danger := roundToCloser(THRESHOLD_DANGER * float64(total) / MEMORY_METRIC_UNIT_AS_BYTES)
	value := roundToCloser(valueInBytes / MEMORY_METRIC_UNIT_AS_BYTES)

	metric := openapi.NewRunnableMetric(
		*openapi.NewNullableString(ptr(MEMORY_METRIC_LABEL)),
		*openapi.NewNullableFloat64(&ratio),
		[]float64{warning, danger},
		*openapi.NewNullableString(ptr(MEMORY_METRIC_UNIT)),
		*openapi.NewNullableFloat64(&value),
	)

	return metric
}

func buildUptimeMetric(uptime time.Duration) *openapi.RunnableMetric {
	value := uptime.Hours()
	unit := "h"
	if value < 1.0 {
		value = uptime.Minutes()
		unit = "min"
		if value < 1.0 {
			value = uptime.Seconds()
			unit = "s"
		}
	}
	value = math.Round(value)

	metric := openapi.NewRunnableMetric(
		*openapi.NewNullableString(ptr(UPTIME_METRIC_LABEL)),
		*openapi.NewNullableFloat64(nil),
		[]float64{},
		*openapi.NewNullableString(&unit),
		*openapi.NewNullableFloat64(&value),
	)

	return metric
}

func checkThatRunnableExists(config *Config, id string) *ServiceError {
	if id != config.runnableServiceSelfId {
		return &ServiceError{HttpStatus: 404, Message: Err404Runnable}
	}
	return nil
}

func getCPUStats() (*uint64, *uint64, error) {
	before, err := cpu.Get()
	if err != nil {
		return nil, nil, err
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		return nil, nil, err
	}

	total := after.Total - before.Total
	value := after.User - before.User

	return &value, &total, nil
}

func getNameFromHostname(config *Config) string {
	name := config.runnableServiceSelfNameFallback
	hostname, err := os.Hostname()
	if err == nil && len(hostname) > 0 {
		name = hostname
	}
	return name
}

func getNullableFrom(value string) *openapi.NullableString {
	if len(value) == 0 {
		return openapi.NewNullableString(nil)
	}

	return openapi.NewNullableString(&value)
}

func getScope(label string, value string) *openapi.NullableRunnableScope {
	if len(label) == 0 || len(value) == 0 {
		return openapi.NewNullableRunnableScope(nil)
	}

	return openapi.NewNullableRunnableScope(
		openapi.NewRunnableScope(
			label,
			value,
		),
	)
}
