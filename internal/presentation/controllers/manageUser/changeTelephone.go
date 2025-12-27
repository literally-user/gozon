package manageUser

import (
	"encoding/json"
	"net/http"

	userApplication "github.com/literally_user/gozon/internal/application/usecases/manageUser"
	"github.com/literally_user/gozon/internal/presentation/controllers/errors"
	"github.com/literally_user/gozon/internal/presentation/middlewares"
)

type ChangeTelephoneRequest struct {
	Telephone string `json:"telephone"`
}

type ChangeTelephoneController struct {
	ChangeTelephoneInteractor userApplication.ChangeTelephoneInteractor
}

func (c *ChangeTelephoneController) Execute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() //nolint:errcheck

	user, ok := r.Context().Value(middlewares.UserContextKey).(middlewares.UserContext)
	if !ok {
		errors.WriteError(w, r, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req ChangeTelephoneRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, "Invalid json body")
		return
	}

	err := c.ChangeTelephoneInteractor.Execute(user.UserUUID, req.Telephone)
	if err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Location", "/users/me/telephone")
	w.WriteHeader(http.StatusNoContent)
}
