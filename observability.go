package main

import (
	"net/http"

	"github.com/nerd500/clinic-appointment/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, "Server Online")
}
