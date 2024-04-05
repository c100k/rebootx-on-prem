package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	config := getConfig()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	dashboardService := loadDashboardService(config)
	runnableService := loadRunnableService(config, logger)

	logger.Info(fmt.Sprintf("Using dashboardServiceImpl : %s", config.dashboardServiceImpl))
	logger.Info(fmt.Sprintf("Using runnableServiceImpl : %s", config.runnableServiceImpl))

	router := mux.NewRouter()

	router.Use(logMiddleware(config, logger))
	router.Use(headerMiddleware(config))
	router.Use(authMiddleware(config))

	rootPath := fmt.Sprintf("/%s", config.pathPrefix)
	dashboardsPath := fmt.Sprintf("%s/dashboards", rootPath)
	runnablesPath := fmt.Sprintf("%s/runnables", rootPath)

	router.HandleFunc(dashboardsPath, getDashboardsHandler(*dashboardService)).Methods("GET")
	router.HandleFunc(runnablesPath, getRunnablesHandler(*runnableService)).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/{id}/reboot", runnablesPath), postRunnableRebootHandler(*runnableService)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/{id}/stop", runnablesPath), postRunnableStopHandler(*runnableService)).Methods("POST")

	headersCORS := handlers.AllowedHeaders([]string{AUTHORIZATION_HEADER, "Content-Type", "Origin"})
	methodsCORS := handlers.AllowedMethods([]string{"GET", "HEAD", "OPTIONS", "POST"})
	originsCORS := handlers.AllowedOrigins([]string{"*"})
	routerWithCORS := handlers.CORS(headersCORS, methodsCORS, originsCORS)(router)

	logger.Info(fmt.Sprintf("Listening on %s://%s:%d", config.protocol, config.bind, config.port))

	http.Handle("/", routerWithCORS)
	http.ListenAndServe(fmt.Sprintf("%s:%d", config.bind, config.port), nil)
}
