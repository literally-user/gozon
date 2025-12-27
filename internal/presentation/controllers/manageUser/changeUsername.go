package manageUser

import (
	"encoding/json"
	"net/http"

	userApplication "github.com/literally_user/gozon/internal/application/usecases/manageUser"
	"github.com/literally_user/gozon/internal/presentation/controllers/errors"
	"github.com/literally_user/gozon/internal/presentation/middlewares"
)

type ChangeUsernameRequest struct {
	Username string `json:"username"`
}

type ChangeUsernameController struct {
	ChangeUsernameInteractor userApplication.ChangeUsernameInteractor
}

func (c *ChangeUsernameController) Execute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() //nolint:errcheck

	user, ok := r.Context().Value(middlewares.UserContextKey).(middlewares.UserContext)
	if !ok {
		errors.WriteError(w, r, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req ChangeUsernameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, "Invalid json body")
		return
	}

	err := c.ChangeUsernameInteractor.Execute(user.UserUUID, req.Username)
	if err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Location", "/users/me/username")
	w.WriteHeader(http.StatusNoContent)
}
