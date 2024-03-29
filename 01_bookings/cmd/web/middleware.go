package main

import (
	"net/http"

	"github.com/Man-Crest/GO-Projects/01_bookings/pkg/helpers"
	"github.com/justinas/nosurf"
)

// NoSurf is the csrf protection middleware
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves session data for current request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a := helpers.IsAuthenticated(r)
		if a != true {
			session.Put(r.Context(), "error", "Please login first")
			http.Redirect(w, r, "/login-user", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
