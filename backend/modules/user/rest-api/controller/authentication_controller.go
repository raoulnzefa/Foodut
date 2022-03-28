package controller

import (
	"encoding/json"
	"net/http"
	"time"

	rspn "github.com/Foodut/backend/responses"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("foodutKey")
var tokenName = "token"

type Claims struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	UserType int    `json:"user_type"`
	jwt.StandardClaims
}

func generateToken(w http.ResponseWriter, id int, email string, userType int) {
	tokenExpiryTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		ID:       id,
		Email:    email,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
	})
}

func resetUserToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})
}

func Authenticate(next http.HandlerFunc, accessType int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidToken := validateUserToken(w, r, accessType)
		if !isValidToken {
			var response rspn.Response
			response.Response_401()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func validateUserToken(w http.ResponseWriter, r *http.Request, accessType int) bool {
	isAccessTokenValid, _, userType := validateTokenFromCookies(r)

	if isAccessTokenValid {
		isUserValid := userType == accessType
		if isUserValid {
			return true
		}
	}
	return false
}

func validateTokenFromCookies(r *http.Request) (bool, string, int) {
	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.Email, accessClaims.UserType
		}
	}
	return false, "", -1
}

func GetEmailFromToken(r *http.Request) string {
	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return accessClaims.Email
		}
	}
	return ""
}
