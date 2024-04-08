package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/YousefAldabbas/auth-service/pkg/model"
	"github.com/YousefAldabbas/auth-service/pkg/repository"
	"github.com/YousefAldabbas/auth-service/pkg/utils"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
)

type UserHandler struct {
	Repo repository.UserRepo
}

func (h UserHandler) GetUserByUUID(w http.ResponseWriter, r *http.Request) {

	userUUID := chi.URLParam(r, "userUUID")

	user, err := h.Repo.GetUserByUUID(userUUID)

	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch user from the database")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Failed to get user",
		})
		return
	}

	if user == (model.User{}) {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "User Doesn't Exist",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": user,
	})

}

func (uh UserHandler) RegisterNewUser(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Error().Err(err).Msg("Error decoding request body")
		utils.ResponseWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user := model.User{
		Username: body.Username,
		UUID:     uuid.New().String(),
		Email:    body.Email,
	}

	hashedPassword, err := utils.HashPassword(body.Password)

	if err != nil {
		log.Error().Err(err).Msg("Unable to hash user's password")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "SERVER ERROR")
		return
	}

	user.Password = hashedPassword

	err = uh.Repo.InsertUser(&user)
	var e *pgconn.PgError

	if err != nil {
		if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
			log.Error().Err(err).Msg("Unique violation error")
			utils.ResponseWithJSON(w, http.StatusConflict, "User already exists")
			return
		}
		log.Error().Err(err).Msg("Unable to insert user to the database")

		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Server error: Unable to insert user to the database.")
		return
	}

	w.WriteHeader(201)
	utils.ResponseWithJSON(w, 201, user)

}
