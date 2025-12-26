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

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginController struct {
	GetUserInteractor userApplication.GetUserInteractor
	TokenManager      auth.TokenManager
}

func (c *LoginController) Execute(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteError(w, r, http.StatusBadRequest, "Invalid json body")
		return
	}

	user, err := c.GetUserInteractor.Execute(req.Username, req.Password)
	if err != nil {
		errors.WriteError(w, r, http.StatusNotFound, err.Error())
		return
	}

	token, err := c.TokenManager.GenerateAuthToken(user.UUID, user.Privileges)
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
	w.Header().Set("Location", "/users/login/"+user.UUID.String())
	w.WriteHeader(http.StatusCreated)
}
