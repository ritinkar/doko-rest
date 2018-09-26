package controllers

import "net/http"

// Root : does root handling
func Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from root \n"))
}
