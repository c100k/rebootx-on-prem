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
	RunnableServiceSelfId                 string
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
		ApiKey:                                getEnvOrPanic("API_KEY"),
		Bind:                                  getEnvOr("BIND", "0.0.0.0"),
		DashboardServiceFileJsonFilePath:      getNullableEnv("DASHBOARD_SERVICE_FILE_JSON_FILE_PATH"),
		DashboardServiceImpl:                  getEnvOr("DASHBOARD_SERVICE_IMPL", "fileJson"),
		PathPrefix:                            getEnvOrPanic("PATH_PREFIX"),
		Port:                                  getEnvAsIntOr("PORT", int32(8080)),
		Protocol:                              getEnvOr("PROTOCOL", "http"),
		RunnableServiceFileJsonFilePath:       getNullableEnv("RUNNABLE_SERVICE_FILE_JSON_FILE_PATH"),
		RunnableServiceImpl:                   getEnvOr("RUNNABLE_SERVICE_IMPL", "fileJson"),
		RunnableServiceSelfFQDN:               getEnvOr("RUNNABLE_SERVICE_SELF_FQDN", ""),
		RunnableServiceSelfFlavor:             getEnvOr("RUNNABLE_SERVICE_SELF_FLAVOR", ""),
		RunnableServiceSelfIPv4:               getEnvOr("RUNNABLE_SERVICE_SELF_IPv4", ""),
		RunnableServiceSelfId:                 getEnvOr("RUNNABLE_SERVICE_SELF_ID", "self"),
		RunnableServiceSelfNameFallback:       getEnvOr("RUNNABLE_SERVICE_SELF_NAME_FALLBACK", "default"),
		RunnableServiceSelfScopesGeoLabel:     getEnvOr("RUNNABLE_SERVICE_SELF_SCOPES_GEO_LABEL", ""),
		RunnableServiceSelfScopesGeoValue:     getEnvOr("RUNNABLE_SERVICE_SELF_SCOPES_GEO_VALUE", ""),
		RunnableServiceSelfScopesLogicalLabel: getEnvOr("RUNNABLE_SERVICE_SELF_SCOPES_LOGICAL_LABEL", ""),
		RunnableServiceSelfScopesLogicalValue: getEnvOr("RUNNABLE_SERVICE_SELF_SCOPES_LOGICAL_VALUE", ""),
		RunnableServiceSelfSSHKeyname:         getEnvOr("RUNNABLE_SERVICE_SELF_SSH_KEYNAME", "default"),
		RunnableServiceSelfSSHPort:            getEnvAsIntOr("RUNNABLE_SERVICE_SELF_SSH_PORT", int32(22)),
		RunnableServiceSelfSSHUsername:        getEnvOr("RUNNABLE_SERVICE_SELF_SSH_USERNAME", "root"),
		RunnableServiceSelfStack:              getEnvOr("RUNNABLE_SERVICE_SELF_STACK", ""),
		RunnableServiceSelfSysCmdPkg:          getEnvOr("RUNNABLE_SERVICE_SYS_CMD_PKG", "syscall"),
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

func envName(key string) string {
	return fmt.Sprintf("%s%s", ENV_VAR_PREFIX, key)
}

func getEnvOr(key string, fallback string) string {
	v := os.Getenv(envName(key))
	if len(v) == 0 {
		return fallback
	}
	return v
}

func getEnvOrPanic(key string) string {
	v := getEnvOr(key, "")
	if len(v) == 0 {
		panic(fmt.Sprintf("You must define the '%s' env var", envName(key)))
	}
	return v
}

func getEnvAsIntOr(key string, fallback int32) int32 {
	raw := getEnvOr(key, "")
	v := utils.ParseInt(&raw)
	if v == nil {
		return fallback
	}
	return *v
}

func getNullableEnv(key string) *string {
	v := getEnvOr(key, "")
	if len(v) == 0 {
		return nil
	}
	return &v
}
