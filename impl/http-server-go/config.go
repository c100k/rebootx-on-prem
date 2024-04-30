package main

import (
	"fmt"
	"os"
	"rebootx-on-prem/http-server-go/utils"
	"slices"
)

type Config struct {
	apiKey                                string
	bind                                  string
	dashboardServiceFileJsonFilePath      *string
	dashboardServiceImpl                  string
	pathPrefix                            string
	port                                  int32
	protocol                              string
	runnableServiceFileJsonFilePath       *string
	runnableServiceImpl                   string
	runnableServiceSelfFQDN               string
	runnableServiceSelfFlavor             string
	runnableServiceSelfIPv4               string
	runnableServiceSelfId                 string
	runnableServiceSelfNameFallback       string
	runnableServiceSelfScopesGeoLabel     string
	runnableServiceSelfScopesGeoValue     string
	runnableServiceSelfScopesLogicalLabel string
	runnableServiceSelfScopesLogicalValue string
	runnableServiceSelfSSHKeyname         string
	runnableServiceSelfSSHPort            int32
	runnableServiceSelfSSHUsername        string
	runnableServiceSelfStack              string
	runnableServiceSelfSysCmdPkg          string
}

const ENV_VAR_PREFIX = "RBTX_"

func getConfig() *Config {
	config := Config{
		apiKey:                                getEnvOrPanic("API_KEY"),
		bind:                                  getEnvOr("BIND", "0.0.0.0"),
		dashboardServiceFileJsonFilePath:      getNullableEnv("DASHBOARD_SERVICE_FILE_JSON_FILE_PATH"),
		dashboardServiceImpl:                  getEnvOr("DASHBOARD_SERVICE_IMPL", "fileJson"),
		pathPrefix:                            getEnvOrPanic("PATH_PREFIX"),
		port:                                  getEnvAsIntOr("PORT", int32(8080)),
		protocol:                              getEnvOr("PROTOCOL", "http"),
		runnableServiceFileJsonFilePath:       getNullableEnv("RUNNABLE_SERVICE_FILE_JSON_FILE_PATH"),
		runnableServiceImpl:                   getEnvOr("RUNNABLE_SERVICE_IMPL", "fileJson"),
		runnableServiceSelfFQDN:               getEnvOr("RUNNABLE_SERVICE_SELF_FQDN", ""),
		runnableServiceSelfFlavor:             getEnvOr("RUNNABLE_SERVICE_SELF_FLAVOR", ""),
		runnableServiceSelfIPv4:               getEnvOr("RUNNABLE_SERVICE_SELF_IPv4", ""),
		runnableServiceSelfId:                 getEnvOr("RUNNABLE_SERVICE_SELF_ID", "self"),
		runnableServiceSelfNameFallback:       getEnvOr("RUNNABLE_SERVICE_SELF_NAME_FALLBACK", "default"),
		runnableServiceSelfScopesGeoLabel:     getEnvOr("RUNNABLE_SERVICE_SELF_SCOPES_GEO_LABEL", ""),
		runnableServiceSelfScopesGeoValue:     getEnvOr("RUNNABLE_SERVICE_SELF_SCOPES_GEO_VALUE", ""),
		runnableServiceSelfScopesLogicalLabel: getEnvOr("RUNNABLE_SERVICE_SELF_SCOPES_LOGICAL_LABEL", ""),
		runnableServiceSelfScopesLogicalValue: getEnvOr("RUNNABLE_SERVICE_SELF_SCOPES_LOGICAL_VALUE", ""),
		runnableServiceSelfSSHKeyname:         getEnvOr("RUNNABLE_SERVICE_SELF_SSH_KEYNAME", "default"),
		runnableServiceSelfSSHPort:            getEnvAsIntOr("RUNNABLE_SERVICE_SELF_SSH_PORT", int32(22)),
		runnableServiceSelfSSHUsername:        getEnvOr("RUNNABLE_SERVICE_SELF_SSH_USERNAME", "root"),
		runnableServiceSelfStack:              getEnvOr("RUNNABLE_SERVICE_SELF_STACK", ""),
		runnableServiceSelfSysCmdPkg:          getEnvOr("RUNNABLE_SERVICE_SYS_CMD_PKG", "syscall"),
	}

	assertOneOf(config.dashboardServiceImpl, []string{"fileJson"})
	assertServiceImplFileJson(config.dashboardServiceImpl, config.dashboardServiceFileJsonFilePath)
	assertOneOf(config.dashboardServiceImpl, []string{"fileJson", "self"})
	assertServiceImplFileJson(config.runnableServiceImpl, config.runnableServiceFileJsonFilePath)
	assertOneOf(config.runnableServiceSelfSysCmdPkg, []string{"exec", "syscall"})

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
