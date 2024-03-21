package main

import (
	"errors"
	"log/slog"
	"openapi"
	"os"
	"os/exec"
	"syscall"
)

type ServiceSelf struct {
	config *Config
	logger *slog.Logger
}

func (service ServiceSelf) list(params *openapi.ListRunnablesQueryParams) (*openapi.ListResRunnable, *ServiceError) {
	config := service.config

	q := params.Q
	if len(*q) > 0 {
		err := checkThatRunnableExists(config, *q)
		if err != nil {
			return openapi.NewListResRunnable([]openapi.Runnable{}, 0), nil
		}
	}

	items := []openapi.Runnable{
		*openapi.NewRunnable(
			*openapi.NewNullableString(&config.runnableFlavor),
			*openapi.NewNullableString(&config.runnableFQDN),
			config.runnableId,
			*openapi.NewNullableString(&config.runnableIPv4),
			nameFromHostname(config),
			*openapi.NewRunnableScopes(
				*openapi.NewNullableRunnableScope(
					openapi.NewRunnableScope(
						config.runnableScopesGeoLabel,
						config.runnableScopesGeoValue,
					),
				),
				*openapi.NewNullableRunnableScope(
					openapi.NewRunnableScope(
						config.runnableScopesLogicalLabel,
						config.runnableScopesLogicalValue,
					),
				),
			),
			*openapi.NewNullableRunnableSSH(
				openapi.NewRunnableSSH(
					*openapi.NewNullableString(&config.runnableSSHKeyname),
					config.runnableSSHPort,
					config.runnableSSHUsername,
				),
			),
			*openapi.NewNullableString(&config.runnableStack),
			openapi.ON,
		),
	}
	total := int32(len(items))

	res := openapi.NewListResRunnable(items, total)

	return res, nil
}

func (service ServiceSelf) reboot(id string) (*openapi.RunnableOperationRes, *ServiceError) {
	config := service.config

	err := checkThatRunnableExists(config, id)
	if err != nil {
		return nil, err
	}

	errExec := execOperation(config, config.sysCmdReboot, syscall.LINUX_REBOOT_CMD_RESTART)
	if errExec != nil {
		return nil, &ServiceError{HttpStatus: 500, Message: errExec.Error()}
	}

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}

func (service ServiceSelf) stop(id string) (*openapi.RunnableOperationRes, *ServiceError) {
	config := service.config

	err := checkThatRunnableExists(config, id)
	if err != nil {
		return nil, err
	}

	errExec := execOperation(config, config.sysCmdStop, syscall.LINUX_REBOOT_CMD_POWER_OFF)
	if errExec != nil {
		return nil, &ServiceError{HttpStatus: 500, Message: errExec.Error()}
	}

	return openapi.NewRunnableOperationRes(*openapi.NewNullableString(nil)), nil
}

func checkThatRunnableExists(config *Config, id string) *ServiceError {
	if id != config.runnableId {
		return &ServiceError{HttpStatus: 404, Message: Err404Runnable}
	}
	return nil
}

func execOperation(config *Config, forExec string, forSyscall int) error {
	switch config.sysCmdPkg {
	case "exec":
		cmd := exec.Command(forExec)
		return cmd.Run()
	case "syscall":
		return syscall.Reboot(forSyscall)
	default:
		return errors.New("Invalid sysCmdPkg")
	}
}

func nameFromHostname(config *Config) string {
	name := config.runnableNameFallback
	hostname, err := os.Hostname()
	if err == nil && len(hostname) > 0 {
		name = hostname
	}
	return name
}
