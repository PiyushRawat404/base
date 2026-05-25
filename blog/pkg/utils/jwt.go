package utils

import (
	"blog/internal/model"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type JWTClaims struct {
	Sub   int        `json:"sub"`
	Email string     `json:"email"`
	Role  model.Role `json:"role"`
	Exp   int64      `json:"exp"`
	Iat   int64      `json:"iat"`
}

func GenerateJWT(secret string, user model.User) (string, error) {
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}

	claims := JWTClaims{
		Sub:   user.ID,
		Email: user.Email,
		Role:  user.Role,
		Exp:   time.Now().Add(24 * time.Hour).Unix(),
		Iat:   time.Now().Unix(),
	}

	return encodeJWT(secret, header, claims)
}

func VerifyJWT(secret, token string) (JWTClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return JWTClaims{}, errors.New("invalid token format")
	}

	signingInput := parts[0] + "." + parts[1]
	expectedSignature := signJWT(secret, signingInput)
	if !hmac.Equal([]byte(expectedSignature), []byte(parts[2])) {
		return JWTClaims{}, errors.New("invalid token signature")
	}

	var claims JWTClaims
	if err := decodeSegment(parts[1], &claims); err != nil {
		return JWTClaims{}, err
	}

	if time.Now().Unix() > claims.Exp {
		return JWTClaims{}, errors.New("token expired")
	}

	return claims, nil
}

func encodeJWT(secret string, header map[string]string, claims JWTClaims) (string, error) {
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	headerSegment := base64.RawURLEncoding.EncodeToString(headerJSON)
	claimsSegment := base64.RawURLEncoding.EncodeToString(claimsJSON)
	signingInput := headerSegment + "." + claimsSegment
	signature := signJWT(secret, signingInput)

	return fmt.Sprintf("%s.%s.%s", headerSegment, claimsSegment, signature), nil
}

func signJWT(secret, payload string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func decodeSegment(segment string, target any) error {
	raw, err := base64.RawURLEncoding.DecodeString(segment)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, target)
}
