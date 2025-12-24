package errors

import (
	"encoding/json"
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
	json.NewEncoder(w).Encode(ErrorDetails{
		Status:   status,
		Detail:   detail,
		Instance: r.URL.Path,
	})
}
