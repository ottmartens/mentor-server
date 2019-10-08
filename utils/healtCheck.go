package utils

import (
	"net/http"
)

var HealthCheck = func(w http.ResponseWriter, r *http.Request) {
	Respond(w, Message(true, "ok"))
}
