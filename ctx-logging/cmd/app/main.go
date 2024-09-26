package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/jibaru/ctx-logging/internal/application"
	"github.com/jibaru/ctx-logging/internal/infrastructure/handlers"
	"github.com/jibaru/ctx-logging/internal/infrastructure/loaders/raw"
	"github.com/jibaru/ctx-logging/internal/infrastructure/logger"
	"github.com/jibaru/ctx-logging/internal/infrastructure/repositories/memory"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_CREATE, os.ModeAppend)
	if err != nil {
		slog.Error("error opening logs.txt", "error", err)
		return
	}
	defer file.Close()

	logHandler := logger.NewLogHandler(slog.NewJSONHandler(file, nil))
	slog.SetDefault(slog.New(logHandler))

	loader := raw.NewLoader("dex.json")
	data, err := loader.Load()
	if err != nil {
		slog.Error("error loading dex.json", "error", err)
		return
	}

	repository := memory.NewRawPokemonRepository(data)
	service := application.NewGetSingleService(repository)

	r := mux.NewRouter()
	r.Handle("/api/pokemon/{id}", handlers.NewGetSingleHandler(service)).Methods(http.MethodGet)

	slog.Info("server started at http://localhost:8080")
	if err = http.ListenAndServe(":8080", r); err != nil {
		slog.Error("error starting server", "error", err)
		return
	}
}
