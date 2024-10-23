package tokenization

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type Tokenization struct {
	secret            string
	expiry            time.Duration
	refreshTokenCache map[string]RefreshTokenModel
}

type RefreshTokenModel struct {
	token  string
	expiry time.Time
}

func NewTokenization() *Tokenization {
	return &Tokenization{
		secret:            viper.GetString("Jwt.Secret"),
		expiry:            viper.GetDuration("Jwt.Expiry"),
		refreshTokenCache: make(map[string]RefreshTokenModel),
	}
}

func (t *Tokenization) ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.secret), nil
	})

	if err != nil {
		log.Printf("An error occurred while validating token: %v", err)
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expiry := int64(claims["exp"].(float64))
		return expiry >= time.Now().Unix(), nil
	}

	return false, errors.New("invalid token")
}

func (t *Tokenization) GenerateToken(id string, role int) (string, error) {
	claims := jwt.MapClaims{
		"id":   id,
		"role": fmt.Sprintf("%d", role),
		"exp":  time.Now().Add(t.expiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(t.secret))
	if err != nil {
		log.Printf("An error occurred while generating token: %v", err)
		return "", err
	}

	return tokenString, nil
}

func (t *Tokenization) GenerateRefreshToken(key string) (string, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		log.Printf("An error occurred while generating refresh token: %v", err)
		return "", err
	}

	refreshToken := base64.StdEncoding.EncodeToString(salt)
	expiry := time.Now().Add(t.expiry)

	t.refreshTokenCache[key] = RefreshTokenModel{
		token:  refreshToken,
		expiry: expiry,
	}

	return refreshToken, nil
}

func (t *Tokenization) VerifyRefreshToken(key, token string) (bool, error) {
	model, exists := t.refreshTokenCache[key]
	if !exists {
		return false, nil
	}

	if model.token != token || model.expiry.Before(time.Now()) {
		return false, nil
	}

	delete(t.refreshTokenCache, key)
	return true, nil
}
