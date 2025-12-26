package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenManager struct {
	SecretKey []byte
}

func (g *TokenManager) GenerateAuthToken(userUUID uuid.UUID, privileges int) (string, error) {
	claims := jwt.MapClaims{
		"user_uuid":  userUUID,
		"privileges": privileges,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(g.SecretKey)
}

func (g *TokenManager) ParseAuthToken(tokenString string) (uuid.UUID, int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(g.SecretKey), nil
	})

	if err != nil {
		return uuid.Nil, 0, err
	}

	if !token.Valid {
		return uuid.Nil, 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, 0, errors.New("invalid claims")
	}

	userUUIDStr, ok := claims["user_uuid"].(string)
	if !ok {
		return uuid.Nil, 0, errors.New("user_uuid not found")
	}
	userUUID, err := uuid.Parse(userUUIDStr)
	if err != nil {
		return uuid.Nil, 0, err
	}

	privilegesFloat, ok := claims["privileges"].(float64)
	if !ok {
		return uuid.Nil, 0, errors.New("privileges not found")
	}
	privileges := int(privilegesFloat)

	return userUUID, privileges, nil
}
