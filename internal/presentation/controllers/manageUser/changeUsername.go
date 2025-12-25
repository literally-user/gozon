package manageUser

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	userApplication "github.com/literally_user/gozon/internal/application/usecases/manageUser"
	"github.com/literally_user/gozon/internal/presentation/controllers/errors"
)

type ChangeUsernameRequest struct {
	UserUUID uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
}

type ChangeUsernameController struct {
	ChangeUsernameInteractor userApplication.ChangeUsernameInteractor
}

func (c *ChangeUsernameController) Execute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() //nolint:errcheck

	var req ChangeUsernameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, "Invalid json body")
		return
	}

	err := c.ChangeUsernameInteractor.Execute(req.UserUUID, req.Username)
	if err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Location", "/users/me/username")
	w.WriteHeader(http.StatusNoContent)
}
