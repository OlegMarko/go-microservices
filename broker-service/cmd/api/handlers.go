package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponce{
		Error:   false,
		Message: "Hit The Broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
