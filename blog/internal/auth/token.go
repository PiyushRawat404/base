package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"blog/internal/model"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

type Claims struct {
	UserID int        `json:"user_id"`
	Email  string     `json:"email"`
	Role   model.Role `json:"role"`
	Exp    int64      `json:"exp"`
}

func GenerateToken(user model.User, secret string, ttl time.Duration) (string, error) {
	claims := Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		Exp:    time.Now().Add(ttl).Unix(),
	}

	payload, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	encodedPayload := base64.RawURLEncoding.EncodeToString(payload)
	signature := sign(encodedPayload, secret)

	return fmt.Sprintf("%s.%s", encodedPayload, signature), nil
}

func ParseToken(token, secret string) (Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return Claims{}, ErrInvalidToken
	}

	if !hmac.Equal([]byte(parts[1]), []byte(sign(parts[0], secret))) {
		return Claims{}, ErrInvalidToken
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return Claims{}, ErrInvalidToken
	}

	var claims Claims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return Claims{}, ErrInvalidToken
	}

	if time.Now().Unix() > claims.Exp {
		return Claims{}, ErrExpiredToken
	}

	return claims, nil
}

func sign(payload, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}
