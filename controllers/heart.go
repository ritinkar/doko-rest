package controllers

import (
	"net/http"
)

// HeartBeat : heartbeat
func HeartBeat(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("still beating \n"))
}
