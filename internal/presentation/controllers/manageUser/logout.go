package manageUser

import (
	"net/http"
)

type LogoutController struct{}

func (c *LogoutController) Execute(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "authentication",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	w.Header().Set("Location", "/users/me/logout")
	w.WriteHeader(http.StatusNoContent)
}
