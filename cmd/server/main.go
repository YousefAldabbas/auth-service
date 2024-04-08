package main

import (
	"context"
	"net/http"

	"github.com/YousefAldabbas/auth-service/internal/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	app := App{}
	app.Start()

}

type App struct {
	Router *chi.Mux
	DB     *pgx.Conn
}

func (a *App) Init() error {
	ctx := context.Background()
	cfg, err := config.Config{}.New()
	if err != nil {
		log.Error().Err(err).Msg("Faild to fetch env variable")
		return err
	}
	conn, err := cfg.NewDBConnection(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Faild to Establish new DB connection")
		return err
	}
	a.DB = conn

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	a.Router = r
	a.LoadRoutes()
	return nil
}


func (a *App) Start() {
	err := a.Init()
	if err != nil {
		log.Error().Err(err).Msg("Failed to initialize application")
		return
	}
	http.ListenAndServe(":8080", a.Router)
}
