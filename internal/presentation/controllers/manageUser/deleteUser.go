package manageUser

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	userApplication "github.com/literally_user/gozon/internal/application/usecases/manageUser"
	"github.com/literally_user/gozon/internal/presentation/controllers/errors"
)

type DeleteUserRequest struct {
	UUID string `json:"UUID"`
}

type DeleteUserController struct {
	DeleteUserInteractor userApplication.DeleteUserInteractor
}

func (c *DeleteUserController) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 2 || parts[0] != "users" {
		errors.WriteError(w, r, http.StatusNotFound, "not found")
		return
	}

	userUUID, err := uuid.Parse(parts[1])
	if err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, "invalid uuid")
		return
	}

	if err := c.DeleteUserInteractor.Execute(userUUID); err != nil {
		errors.WriteError(w, r, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
