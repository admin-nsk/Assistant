package jwt

import (
	"context"
	"encoding/json"
	"net/http"
	"note_service/app/pkg/logging"
	"note_service/app/internal/app_context"
	"strings"
	"time"

	"github.com/cristalhq/jwt/v3"
)


type UserClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

func JWTMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer")
		logger := logging.Getlogger()
		if len(authHeader) != 2 {
			logger.Error("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("malformed token"))
			return
		}
		logger.Debug("create JWT verifier")
		jwtToken := strings.TrimSpace(authHeader[1])
		key := []byte(app_context.GetInstance().Config.JWT.Secret)
		logger.Debug("key ", key)
		verifier, err := jwt.NewVerifierHS(jwt.HS256, key)
		if err != nil {
			unauthorized(w, err)
			return
		}
		logger.Debug("parse and verify token")
		token, err := jwt.ParseAndVerifyString(jwtToken, verifier)
		if err != nil {
			unauthorized(w, err)
			return
		}
		logger.Debug("parse user claims")
		var uc UserClaims
		err = json.Unmarshal(token.RawClaims(), &uc)
		if err != nil {
			unauthorized(w, err)
			return
		}
		if valid := uc.IsValidAt(time.Now()); !valid {
			logging.Getlogger().Error("token has been expired")
			unauthorized(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", uc.ID)
		h(w, r.WithContext(ctx))

	}
}


func unauthorized(w http.ResponseWriter, err error)  {
	logging.Getlogger().Error(err)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("unauthorized"))
}