package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gym-genius/internal/config"
	"gym-genius/internal/http-server/handlers/traning/save"
	"gym-genius/internal/storage"
	"gym-genius/internal/storage/trainings"
	"net/http"
)

func main() {
	cfg := config.MustLoadConfig()
	db := storage.New(cfg.DB.Url)

	trainingsStorage := trainings.New(db)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", save.New(trainingsStorage))

	srv := &http.Server{
		Addr:         cfg.Http.Address,
		Handler:      router,
		WriteTimeout: 15 * 60, // 15 minutes
		ReadTimeout:  15 * 60, // 15 minutes
	}

	if err := srv.ListenAndServe(); err != nil {
		panic("failed to start server: " + err.Error())
	}
}
