package manageUser

import (
	"net/http"

	userApplication "github.com/literally_user/gozon/internal/application/usecases/manageUser"
	"github.com/literally_user/gozon/internal/presentation/controllers/errors"
	"github.com/literally_user/gozon/internal/presentation/middlewares"
)

type DeleteUserController struct {
	DeleteUserInteractor userApplication.DeleteUserInteractor
}

func (c *DeleteUserController) Execute(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middlewares.UserContextKey).(middlewares.UserContext)
	if !ok {
		errors.WriteError(w, r, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := c.DeleteUserInteractor.Execute(user.UserUUID); err != nil {
		errors.WriteError(w, r, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Location", "/users/me")
	w.WriteHeader(http.StatusNoContent)
}
