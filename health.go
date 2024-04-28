package jango

import "net/http"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	JsonResponse(w, " server is healthy", 200)
}
