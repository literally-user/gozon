package manageUser

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	userApplication "github.com/literally_user/gozon/internal/application/usecases/manageUser"
	"github.com/literally_user/gozon/internal/presentation/controllers/errors"
)

type ChangeEmailRequest struct {
	UserUUID uuid.UUID `json:"uuid"`
	Email    string    `json:"email"`
}

type ChangeEmailController struct {
	ChangeEmailInteractor userApplication.ChangeEmailInteractor
}

func (c *ChangeEmailController) Execute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() //nolint:errcheck

	var req ChangeEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, "Invalid json body")
		return
	}

	err := c.ChangeEmailInteractor.Execute(req.UserUUID, req.Email)
	if err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Location", "/users/me/email")
	w.WriteHeader(http.StatusNoContent)
}
