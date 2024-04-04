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

	var runnableService RunnableService
	switch config.runnableServiceImpl {
	case "fileJson":
		runnableService = RunnableServiceFileJson{config: config, logger: logger}
	case "noop":
		runnableService = RunnableServiceNoop{logger: logger}
	case "self":
		runnableService = RunnableServiceSelf{config: config, logger: logger}
	default:
		panic(fmt.Sprintf("Invalid runnableServiceImpl : %s", config.runnableServiceImpl))
	}

	logger.Info(fmt.Sprintf("Using runnableServiceImpl : %s", config.runnableServiceImpl))

	router := mux.NewRouter()

	router.Use(logMiddleware(config, logger))
	router.Use(headerMiddleware(config))
	router.Use(authMiddleware(config))

	rootPath := fmt.Sprintf("/%s/runnables", config.pathPrefix)

	router.HandleFunc(rootPath, getRunnablesHandler(runnableService)).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/{id}/reboot", rootPath), postRunnableRebootHandler(runnableService)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/{id}/stop", rootPath), postRunnableStopHandler(runnableService)).Methods("POST")

	headersCORS := handlers.AllowedHeaders([]string{AUTHORIZATION_HEADER, "Content-Type", "Origin"})
	methodsCORS := handlers.AllowedMethods([]string{"GET", "HEAD", "OPTIONS", "POST"})
	originsCORS := handlers.AllowedOrigins([]string{"*"})
	routerWithCORS := handlers.CORS(headersCORS, methodsCORS, originsCORS)(router)

	logger.Info(fmt.Sprintf("Listening on %s://%s:%d", config.protocol, config.bind, config.port))

	http.Handle("/", routerWithCORS)
	http.ListenAndServe(fmt.Sprintf("%s:%d", config.bind, config.port), nil)
}
