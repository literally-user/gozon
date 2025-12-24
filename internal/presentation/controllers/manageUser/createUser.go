package manageUser

import (
	"encoding/json"
	"net/http"

	userApplication "github.com/literally_user/gozon/internal/application/usecases/manageUser"
	"github.com/literally_user/gozon/internal/presentation/controllers/errors"
)

type CreateUserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}

type CreateUserController struct {
	CreateUserInteractor userApplication.CreateUserInteractor
}

func (c *CreateUserController) Execute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, "Invalid json body")
		return
	}

	newUser, err := c.CreateUserInteractor.Execute(userApplication.DTO(req))
	if err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Location", "/users/"+newUser.UUID.String())
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
