package domain

import (
	"bankingApp/domain/logger"
	"encoding/json"
	"net/http"
	"net/url"
)

type AuthRepository interface {
	IsAuthorized(token, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct{}

func (r RemoteAuthRepository) IsAuthorized(token, routeName string, vars map[string]string) bool {
	u := buildVerifyUrl(token, routeName, vars)
	if res, err := http.Get(u); err != nil {
		logger.Debug("Error while sending " + err.Error())
		return false
	} else {
		m := map[string]bool{}
		if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding res from auth server " + err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

// sample url for verify
// /auth/verify?token=a.b.c&routeName=MakeTransaction&customer_id=2002&account_id=1231
func buildVerifyUrl(token, routeName string, vars map[string]string) string {
	u := url.URL{Host: "localhost:8081", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}
