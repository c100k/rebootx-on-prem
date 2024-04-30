package config

import (
	"fmt"
	"os"
	"rebootx-on-prem/http-server-go/utils"
	"slices"
)

type Config struct {
	ApiKey                                string
	Bind                                  string
	DashboardServiceFileJsonFilePath      *string
	DashboardServiceImpl                  string
	PathPrefix                            string
	Port                                  int32
	Protocol                              string
	RunnableServiceFileJsonFilePath       *string
	RunnableServiceImpl                   string
	RunnableServiceSelfFQDN               string
	RunnableServiceSelfFlavor             string
	RunnableServiceSelfIPv4               string
	RunnableServiceSelfID                 string
	RunnableServiceSelfNameFallback       string
	RunnableServiceSelfScopesGeoLabel     string
	RunnableServiceSelfScopesGeoValue     string
	RunnableServiceSelfScopesLogicalLabel string
	RunnableServiceSelfScopesLogicalValue string
	RunnableServiceSelfSSHKeyname         string
	RunnableServiceSelfSSHPort            int32
	RunnableServiceSelfSSHUsername        string
	RunnableServiceSelfStack              string
	RunnableServiceSelfSysCmdPkg          string
}

const ENV_VAR_PREFIX = "RBTX_"

func GetConfig() *Config {
	config := Config{
		ApiKey:                                envOrPanic("API_KEY"),
		Bind:                                  envOr("BIND", "0.0.0.0"),
		DashboardServiceFileJsonFilePath:      nullableEnv("DASHBOARD_SERVICE_FILE_JSON_FILE_PATH"),
		DashboardServiceImpl:                  envOr("DASHBOARD_SERVICE_IMPL", "fileJson"),
		PathPrefix:                            envOrPanic("PATH_PREFIX"),
		Port:                                  envAsIntOr("PORT", int32(8080)),
		Protocol:                              envOr("PROTOCOL", "http"),
		RunnableServiceFileJsonFilePath:       nullableEnv("RUNNABLE_SERVICE_FILE_JSON_FILE_PATH"),
		RunnableServiceImpl:                   envOr("RUNNABLE_SERVICE_IMPL", "fileJson"),
		RunnableServiceSelfFQDN:               envOr("RUNNABLE_SERVICE_SELF_FQDN", ""),
		RunnableServiceSelfFlavor:             envOr("RUNNABLE_SERVICE_SELF_FLAVOR", ""),
		RunnableServiceSelfIPv4:               envOr("RUNNABLE_SERVICE_SELF_IPv4", ""),
		RunnableServiceSelfID:                 envOr("RUNNABLE_SERVICE_SELF_ID", "self"),
		RunnableServiceSelfNameFallback:       envOr("RUNNABLE_SERVICE_SELF_NAME_FALLBACK", "default"),
		RunnableServiceSelfScopesGeoLabel:     envOr("RUNNABLE_SERVICE_SELF_SCOPES_GEO_LABEL", ""),
		RunnableServiceSelfScopesGeoValue:     envOr("RUNNABLE_SERVICE_SELF_SCOPES_GEO_VALUE", ""),
		RunnableServiceSelfScopesLogicalLabel: envOr("RUNNABLE_SERVICE_SELF_SCOPES_LOGICAL_LABEL", ""),
		RunnableServiceSelfScopesLogicalValue: envOr("RUNNABLE_SERVICE_SELF_SCOPES_LOGICAL_VALUE", ""),
		RunnableServiceSelfSSHKeyname:         envOr("RUNNABLE_SERVICE_SELF_SSH_KEYNAME", "default"),
		RunnableServiceSelfSSHPort:            envAsIntOr("RUNNABLE_SERVICE_SELF_SSH_PORT", int32(22)),
		RunnableServiceSelfSSHUsername:        envOr("RUNNABLE_SERVICE_SELF_SSH_USERNAME", "root"),
		RunnableServiceSelfStack:              envOr("RUNNABLE_SERVICE_SELF_STACK", ""),
		RunnableServiceSelfSysCmdPkg:          envOr("RUNNABLE_SERVICE_SYS_CMD_PKG", "syscall"),
	}

	assertOneOf(config.DashboardServiceImpl, []string{"fileJson"})
	assertServiceImplFileJson(config.DashboardServiceImpl, config.DashboardServiceFileJsonFilePath)
	assertOneOf(config.DashboardServiceImpl, []string{"fileJson", "self"})
	assertServiceImplFileJson(config.RunnableServiceImpl, config.RunnableServiceFileJsonFilePath)
	assertOneOf(config.RunnableServiceSelfSysCmdPkg, []string{"exec", "syscall"})

	return &config
}

func assertOneOf(value string, expected []string) {
	idx := slices.Index(expected, value)
	if idx == -1 {
		panic(fmt.Sprintf("Valid values are : %s. Got '%s'", expected, value))
	}
}

func assertServiceImplFileJson(serviceImpl string, filePath *string) {
	if serviceImpl != "fileJson" {
		return
	}

	if filePath == nil {
		panic("You must provide a json file path when using 'fileJson'")
	}

	path := *filePath
	_, err := os.Stat(path)
	if err != nil {
		panic(fmt.Sprintf("The file %s does not exist", path))
	}
}

func envAsIntOr(key string, fallback int32) int32 {
	raw := envOr(key, "")
	v := utils.ParseInt(&raw)
	if v == nil {
		return fallback
	}
	return *v
}

func envName(key string) string {
	return fmt.Sprintf("%s%s", ENV_VAR_PREFIX, key)
}

func envOr(key string, fallback string) string {
	v := os.Getenv(envName(key))
	if len(v) == 0 {
		return fallback
	}
	return v
}

func envOrPanic(key string) string {
	v := envOr(key, "")
	if len(v) == 0 {
		panic(fmt.Sprintf("You must define the '%s' env var", envName(key)))
	}
	return v
}

func nullableEnv(key string) *string {
	v := envOr(key, "")
	if len(v) == 0 {
		return nil
	}
	return &v
}
