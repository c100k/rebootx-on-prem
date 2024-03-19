package main

import (
	"fmt"
	"os"
)

type Config struct {
	apiKey                     string
	bind                       string
	pathPrefix                 string
	port                       int32
	provider                   string
	protocol                   string
	runnableFlavor             string
	runnableFQDN               string
	runnableGroupIdentifier    string
	runnableId                 string
	runnableIPv4               string
	runnableKind               string
	runnableNameFallback       string
	runnableScopesGeoLabel     string
	runnableScopesGeoValue     string
	runnableScopesLogicalLabel string
	runnableScopesLogicalValue string
	runnableSSHKeyname         string
	runnableSSHPort            int32
	runnableSSHUsername        string
	runnableStack              string
	serviceImpl                string
	sysCmdPkg                  string
	sysCmdReboot               string
	sysCmdStop                 string
}

const ENV_VAR_PREFIX = "RBTX_"

func getConfig() *Config {
	config := Config{
		apiKey:                     getEnvOrPanic("API_KEY"),
		bind:                       getEnvOr("BIND", "0.0.0.0"),
		pathPrefix:                 getEnvOrPanic("PATH_PREFIX"),
		port:                       getEnvAsIntOr("PORT", int32(8080)),
		protocol:                   getEnvOr("PROTOCOL", "http"),
		runnableFlavor:             getEnvOr("RUNNABLE_FLAVOR", ""),
		runnableFQDN:               getEnvOr("RUNNABLE_FQDN", ""),
		runnableId:                 getEnvOr("RUNNABLE_ID", "self"),
		runnableIPv4:               getEnvOr("RUNNABLE_IPv4", ""),
		runnableNameFallback:       getEnvOr("RUNNABLE_NAME_FALLBACK", "default"),
		runnableScopesGeoLabel:     getEnvOr("RUNNABLE_SCOPES_GEO_LABEL", "World"),
		runnableScopesGeoValue:     getEnvOr("RUNNABLE_SCOPES_GEO_LABEL", "world"),
		runnableScopesLogicalLabel: getEnvOr("RUNNABLE_SCOPES_LOGICAL_LABEL", "Project 01"),
		runnableScopesLogicalValue: getEnvOr("RUNNABLE_SCOPES_LOGICAL_LABEL", "project-01"),
		runnableSSHKeyname:         getEnvOr("RUNNABLE_SSH_KEYNAME", "default"),
		runnableSSHPort:            getEnvAsIntOr("RUNNABLE_SSH_PORT", int32(22)),
		runnableSSHUsername:        getEnvOr("RUNNABLE_SSH_USERNAME", "root"),
		runnableStack:              getEnvOr("RUNNABLE_STACK", "nodejs"),
		serviceImpl:                getEnvOr("SERVICE_IMPL", "self"),
		sysCmdPkg:                  getEnvOr("SYS_CMD_PKG", "syscall"),
		sysCmdReboot:               getEnvOr("SYS_CMD_REBOOT", "reboot"),
		sysCmdStop:                 getEnvOr("SYS_CMD_STOP", "shutdown"),
	}

	if config.serviceImpl != "noop" && config.serviceImpl != "self" {
		panic(fmt.Sprintf("Valid values for serviceImpl are : 'noop' and 'self'. Got '%s'", config.serviceImpl))
	}

	if config.sysCmdPkg != "exec" && config.sysCmdPkg != "syscall" {
		panic(fmt.Sprintf("Valid values for sysCmdPkg are : 'exec' and 'syscall'. Got '%s'", config.sysCmdPkg))
	}

	return &config
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
	v := parseInt(&raw)
	if v == nil {
		return fallback
	}
	return *v
}
