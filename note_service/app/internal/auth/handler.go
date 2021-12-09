package auth

import (
	"encoding/json"
	"net/http"
	"note_service/app/internal/app_context"
	"note_service/app/pkg/logging"
	jwt2 "note_service/app/pkg/middleware/jwt"
	"time"

	"github.com/cristalhq/jwt/v3"
)

const (
	URL = "/api/auth"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth(w http.ResponseWriter, r *http.Request)  {
	var u user
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		logging.Getlogger().Error(r.Body)
		logging.Getlogger().Fatal(err)
	}

	defer r.Body.Close()

	//TODO

	if u.Username != "me" || u.Password != "password" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	key := []byte(app_context.GetInstance().Config.JWT.Secret)
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		w.WriteHeader(418)
		return
	}

	builder := jwt.NewBuilder(signer)


	//TODO Insert real user data
	claims := jwt2.UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID: "uid_here",
			Audience: []string{"users"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 35)),
		},
		Email: "email",
	}
	token, err := builder.Build(claims)
	if err != nil {
		logging.Getlogger().Error(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	jsonBytes, err := json.Marshal(map[string]string{
		"token": token.String(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	w.WriteHeader(200)
}