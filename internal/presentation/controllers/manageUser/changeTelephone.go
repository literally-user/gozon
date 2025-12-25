package manageUser

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	userApplication "github.com/literally_user/gozon/internal/application/usecases/manageUser"
	"github.com/literally_user/gozon/internal/presentation/controllers/errors"
)

type ChangeTelephoneRequest struct {
	UserUUID  uuid.UUID `json:"uuid"`
	Telephone string    `json:"telephone"`
}

type ChangeTelephoneController struct {
	ChangeTelephoneInteractor userApplication.ChangeTelephoneInteractor
}

func (c *ChangeTelephoneController) Execute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() //nolint:errcheck

	var req ChangeTelephoneRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, "Invalid json body")
		return
	}

	err := c.ChangeTelephoneInteractor.Execute(req.UserUUID, req.Telephone)
	if err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Location", "/users/me/telephone")
	w.WriteHeader(http.StatusNoContent)
}
