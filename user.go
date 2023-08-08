package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nerd500/clinic-appointment/internal/database"
	"github.com/nerd500/clinic-appointment/utils"
)

func (apiCfg *apiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		ContactNo string `json:"contact_no"`
		Password  string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s ", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		FirstName:      params.FirstName,
		LastName:       params.LastName,
		ContactNo:      params.ContactNo,
		UserKey:        utils.GenerateUserKey(params.FirstName, params.LastName, params.ContactNo, params.Password),
		IsAdmin:        false,
		HashedPassword: utils.Hash(params.Password),
	})

	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Failed to create user, Internal server error: %v", err))
		return
	}
	utils.RespondWithJSON(w, 200, user)

}
