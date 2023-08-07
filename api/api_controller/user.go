package controller

import (
	"net/http"

	"github.com/nerd500/clinic-appointment/utils"
)

func userLogin(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, "Server Online")
}
