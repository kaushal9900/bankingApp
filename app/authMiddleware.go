package app

import (
	"bankingApp/domain"
	"bankingApp/errs"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (a AuthMiddleware) authorizedHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")
			if authHeader != "" {
				token := getTokenFromHeader(authHeader)
				isAuthorized := a.repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)
				if isAuthorized {
					next.ServeHTTP(w, r)
				} else {
					appErr := errs.AppError{http.StatusForbidden, "Unauthorized"}
					writeResponse(w, appErr.Code, appErr.AsMessage())
				}
			} else {
				writeResponse(w, http.StatusUnauthorized, "Missing Token")
			}
		})
	}
}

func getTokenFromHeader(token string) string {
	/*
	   token is coming in the format as below
	   "Bearer a.b.c"
	*/
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
