package controllers

import (
	"doko-rest/utils"
	"net/http"
)

// HeartBeat : heartbeat
func HeartBeat(w http.ResponseWriter, r *http.Request) {
	res := utils.Message(true, "Still Beating")
	utils.Respond(w, res)
}
