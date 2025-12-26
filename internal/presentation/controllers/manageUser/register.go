package manageUser

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	userApplication "github.com/literally_user/gozon/internal/application/usecases/manageUser"
	"github.com/literally_user/gozon/internal/infrastructure/auth"
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
	TokenManager         auth.TokenManager
}

func (c *CreateUserController) Execute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() //nolint:errcheck

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

	token, err := c.TokenManager.GenerateAuthToken(newUser.UUID, newUser.Privileges)
	if err != nil {
		log.Printf("Failed to generate auth token: %v", err)
	}

	authCookie := http.Cookie{
		Name:    "authentication",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
		MaxAge:  24 * 60 * 60,

		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &authCookie)
	w.Header().Set("Location", "/users/register")
	w.WriteHeader(http.StatusCreated)
}
