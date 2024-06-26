package self

import (
	"log/slog"
	"math"
	"openapi"
	"os"
	"rebootx-on-prem/http-server-go/config"
	"rebootx-on-prem/http-server-go/resources/runnable/commons"
	"rebootx-on-prem/http-server-go/utils"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/uptime"
)

type Service struct {
	config *config.Config
	logger *slog.Logger
}

func NewService(config *config.Config, logger *slog.Logger) *Service {
	return &Service{config: config, logger: logger}
}

const CPU_METRIC_LABEL = "CPU"
const CPU_METRIC_UNIT = "%"
const MEMORY_METRIC_LABEL = "RAM"
const MEMORY_METRIC_UNIT = "MB"
const MEMORY_METRIC_UNIT_AS_BYTES = 10e6
const THRESHOLD_WARNING = 0.75
const THRESHOLD_DANGER = 0.85
const UPTIME_METRIC_LABEL = "Uptime"

func (service Service) List(params *openapi.ListRunnablesQueryParams) (*openapi.ListResRunnable, *utils.ServiceError) {
	config := service.config

	q := params.Q
	if q != nil && len(*q) > 0 {
		err := checkThatRunnableExists(config, *q)
		if err != nil {
			return openapi.NewListResRunnable([]openapi.Runnable{}, 0), nil
		}
	}

	metrics := []openapi.RunnableMetric{}

	cpuValue, cpuTotal, err := cpuStats()
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
			*nullableFrom(config.RunnableServiceSelfFlavor),
			*nullableFrom(config.RunnableServiceSelfFQDN),
			config.RunnableServiceSelfID,
			*nullableFrom(config.RunnableServiceSelfIPv4),
			metrics,
			nameFromHostname(config),
			*openapi.NewRunnableScopes(
				*scope(config.RunnableServiceSelfScopesGeoLabel, config.RunnableServiceSelfScopesGeoValue),
				*scope(config.RunnableServiceSelfScopesLogicalLabel, config.RunnableServiceSelfScopesLogicalValue),
			),
			*openapi.NewNullableRunnableSSH(
				openapi.NewRunnableSSH(
					*openapi.NewNullableString(&config.RunnableServiceSelfSSHKeyname),
					config.RunnableServiceSelfSSHPort,
					config.RunnableServiceSelfSSHUsername,
				),
			),
			*nullableFrom(config.RunnableServiceSelfStack),
			openapi.ON,
		),
	}
	total := int32(len(items))

	res := openapi.NewListResRunnable(items, total)

	return res, nil
}

func (service Service) Reboot(id string) (*openapi.RunnableOperationRes, *utils.ServiceError) {
	config := service.config

	err := checkThatRunnableExists(config, id)
	if err != nil {
		return nil, err
	}

	errExec := performOpOnSelf(config, commons.REBOOT)
	if errExec != nil {
		return nil, &utils.ServiceError{HTTPStatus: 500, Message: errExec.Error()}
	}

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}

func (service Service) Stop(id string) (*openapi.RunnableOperationRes, *utils.ServiceError) {
	config := service.config

	err := checkThatRunnableExists(config, id)
	if err != nil {
		return nil, err
	}

	errExec := performOpOnSelf(config, commons.STOP)
	if errExec != nil {
		return nil, &utils.ServiceError{HTTPStatus: 500, Message: errExec.Error()}
	}

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}

func buildCPUMetric(used uint64, total uint64) *openapi.RunnableMetric {
	value := utils.RoundToCloser(float64(used) / float64(total) * 100)

	metric := openapi.NewRunnableMetric(
		*openapi.NewNullableString(utils.Ptr(CPU_METRIC_LABEL)),
		*openapi.NewNullableFloat64(nil),
		[]float64{THRESHOLD_WARNING, THRESHOLD_DANGER},
		*openapi.NewNullableString(utils.Ptr(CPU_METRIC_UNIT)),
		*openapi.NewNullableFloat64(&value),
	)

	return metric
}

func buildMemoryMetric(used uint64, total uint64) *openapi.RunnableMetric {
	valueInBytes := float64(used)

	ratio := utils.RoundToCloser(valueInBytes / float64(total))
	warning := utils.RoundToCloser(THRESHOLD_WARNING * float64(total) / MEMORY_METRIC_UNIT_AS_BYTES)
	danger := utils.RoundToCloser(THRESHOLD_DANGER * float64(total) / MEMORY_METRIC_UNIT_AS_BYTES)
	value := utils.RoundToCloser(valueInBytes / MEMORY_METRIC_UNIT_AS_BYTES)

	metric := openapi.NewRunnableMetric(
		*openapi.NewNullableString(utils.Ptr(MEMORY_METRIC_LABEL)),
		*openapi.NewNullableFloat64(&ratio),
		[]float64{warning, danger},
		*openapi.NewNullableString(utils.Ptr(MEMORY_METRIC_UNIT)),
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
		*openapi.NewNullableString(utils.Ptr(UPTIME_METRIC_LABEL)),
		*openapi.NewNullableFloat64(nil),
		[]float64{},
		*openapi.NewNullableString(&unit),
		*openapi.NewNullableFloat64(&value),
	)

	return metric
}

func checkThatRunnableExists(config *config.Config, id string) *utils.ServiceError {
	if id != config.RunnableServiceSelfID {
		return &utils.ServiceError{HTTPStatus: 404, Message: utils.Err404}
	}
	return nil
}

func cpuStats() (*uint64, *uint64, error) {
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

func nameFromHostname(config *config.Config) string {
	name := config.RunnableServiceSelfNameFallback
	hostname, err := os.Hostname()
	if err == nil && len(hostname) > 0 {
		name = hostname
	}
	return name
}

func nullableFrom(value string) *openapi.NullableString {
	if len(value) == 0 {
		return openapi.NewNullableString(nil)
	}

	return openapi.NewNullableString(&value)
}

func scope(label string, value string) *openapi.NullableRunnableScope {
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
