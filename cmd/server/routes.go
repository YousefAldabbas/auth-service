package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/YousefAldabbas/auth-service/pkg/handler"
	"github.com/YousefAldabbas/auth-service/pkg/repository"
	"github.com/go-chi/chi"
)

func (a *App) LoadRoutes() {

	a.Router.Route("/users", func(r chi.Router) {
		uh := handler.UserHandler{Repo: repository.UserRepo{
			DB: a.DB,
		}}
		r.Get("/{userUUID}", uh.GetUserByUUID)
		r.Post("/", uh.RegisterNewUser)
	})

	a.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		err := a.DB.Ping(ctx)

		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode("Exception accure when ping the database")
			return
		}


		w.WriteHeader(200)


		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "OK",
		})
	})

}
