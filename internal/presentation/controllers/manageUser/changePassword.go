package manageUser

import (
	"encoding/json"
	"net/http"

	userApplication "github.com/literally_user/gozon/internal/application/usecases/manageUser"
	"github.com/literally_user/gozon/internal/presentation/controllers/errors"
	"github.com/literally_user/gozon/internal/presentation/middlewares"
)

type ChangePasswordRequest struct {
	Password string `json:"password"`
}

type ChangePasswordController struct {
	ChangePasswordInteractor userApplication.ChangePasswordInteractor
}

func (c *ChangePasswordController) Execute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() //nolint:errcheck

	user, ok := r.Context().Value(middlewares.UserContextKey).(middlewares.UserContext)
	if !ok {
		errors.WriteError(w, r, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req ChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, "Invalid json body")
		return
	}

	err := c.ChangePasswordInteractor.Execute(user.UserUUID, req.Password)
	if err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Location", "/users/me/password")
	w.WriteHeader(http.StatusNoContent)
}
