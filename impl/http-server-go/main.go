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

	var service Service
	switch config.serviceImpl {
	case "fileJson":
		service = ServiceFileJson{config: config, logger: logger}
	case "noop":
		service = ServiceNoop{logger: logger}
	case "self":
		service = ServiceSelf{config: config, logger: logger}
	default:
		panic(fmt.Sprintf("Invalid serviceImpl : %s", config.serviceImpl))
	}

	logger.Info(fmt.Sprintf("Using serviceImpl : %s", config.serviceImpl))

	router := mux.NewRouter()

	router.Use(logMiddleware(config, logger))
	router.Use(headerMiddleware(config))
	router.Use(authMiddleware(config))

	rootPath := fmt.Sprintf("/%s/runnables", config.pathPrefix)

	router.HandleFunc(rootPath, getRunnablesHandler(service)).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/reboot/{id}", rootPath), postRunnableRebootHandler(service)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/stop/{id}", rootPath), postRunnableStopHandler(service)).Methods("POST")

	headersCORS := handlers.AllowedHeaders([]string{AUTHORIZATION_HEADER, "Content-Type", "Origin"})
	methodsCORS := handlers.AllowedMethods([]string{"GET", "HEAD", "OPTIONS", "POST"})
	originsCORS := handlers.AllowedOrigins([]string{"*"})
	routerWithCORS := handlers.CORS(headersCORS, methodsCORS, originsCORS)(router)

	logger.Info(fmt.Sprintf("Listening on %s://%s:%d", config.protocol, config.bind, config.port))

	http.Handle("/", routerWithCORS)
	http.ListenAndServe(fmt.Sprintf("%s:%d", config.bind, config.port), nil)
}
