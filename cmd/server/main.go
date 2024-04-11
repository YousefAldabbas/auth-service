package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/YousefAldabbas/auth-service/internal/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	// "github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	app := App{}
	app.Start(ctx)

}

type App struct {
	Router *chi.Mux
	DB     *sql.DB
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

func (a *App) Start(ctx context.Context) error {
	err := a.Init()
	if err != nil {
		log.Error().Err(err).Msg("Failed to initialize application")
		return err
	}

	defer func(){
		a.DB.Close()
	}()

	ch := make(chan error, 1)

	server := &http.Server{
		Addr:    ":8080",
		Handler: a.Router,
	}
	go func() {
		err = server.ListenAndServe()

		if err != nil {
			ch <- fmt.Errorf("failed to start server :%w", err)
		}

		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
