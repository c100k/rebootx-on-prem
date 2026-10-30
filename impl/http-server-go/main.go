package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"rebootx-on-prem/http-server-go/config"
	"rebootx-on-prem/http-server-go/resources/dashboard"
	"rebootx-on-prem/http-server-go/resources/runnable"

	"github.com/gorilla/handlers"
)

func main() {
	config := config.GetConfig()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	dashboardService := dashboard.LoadService(config)
	runnableService := runnable.LoadService(config, logger)

	logger.Info(fmt.Sprintf("Using DashboardServiceImpl : %s", config.DashboardServiceImpl))
	logger.Info(fmt.Sprintf("Using RunnableServiceImpl : %s", config.RunnableServiceImpl))

	mux := http.NewServeMux()

	rootPath := fmt.Sprintf("/%s", config.PathPrefix)
	dashboardsPath := fmt.Sprintf("%s/dashboards", rootPath)
	runnablesPath := fmt.Sprintf("%s/runnables", rootPath)

	mux.HandleFunc("GET "+dashboardsPath, getDashboardsHandler(*dashboardService))
	mux.HandleFunc("GET "+runnablesPath, getRunnablesHandler(*runnableService))
	mux.HandleFunc("POST "+fmt.Sprintf("%s/{id}/reboot", runnablesPath), postRunnableRebootHandler(*runnableService))
	mux.HandleFunc("POST "+fmt.Sprintf("%s/{id}/stop", runnablesPath), postRunnableStopHandler(*runnableService))

	handler := authMiddleware(config)(mux)
	handler = headerMiddleware(config)(handler)
	handler = logMiddleware(config, logger)(handler)

	headersCORS := handlers.AllowedHeaders([]string{AUTHORIZATION_HEADER, "Content-Type", "Origin"})
	methodsCORS := handlers.AllowedMethods([]string{"GET", "HEAD", "OPTIONS", "POST"})
	originsCORS := handlers.AllowedOrigins([]string{"*"})
	handler = handlers.CORS(headersCORS, methodsCORS, originsCORS)(handler)

	logger.Info(fmt.Sprintf("Listening on %s://%s:%d", config.Protocol, config.Bind, config.Port))

	http.ListenAndServe(fmt.Sprintf("%s:%d", config.Bind, config.Port), handler)
}
