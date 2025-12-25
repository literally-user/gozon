package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorDetails struct {
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

func WriteError(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	detail string,
) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(ErrorDetails{
		Status:   status,
		Detail:   detail,
		Instance: r.URL.Path,
	})
	if err != nil {
		log.Printf("Error writer: %v", err)
	}
}
