package token

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type TokenService interface {
	GenerateToken(userID uuid.UUID) (string, time.Time, error)
	ValidateToken(token string) (uuid.UUID, error)
}

type tokenService struct {
	secretKey string
}

func NewTokenServiceImpl(secretKey string) *tokenService {
	return &tokenService{secretKey: secretKey}
}

func (t *tokenService) GenerateToken(userID uuid.UUID) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(t.secretKey))
	if err != nil {
		return "", time.Time{}, err
	}
	return signedToken, expirationTime, nil
}

func (t *tokenService) ValidateToken(token string) (uuid.UUID, error) {
	claims := &jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secretKey), nil
	})
	if err != nil || !parsedToken.Valid {
		return uuid.Nil, err
	}

	userIDStr, ok := (*claims)["user_id"].(string)
	if !ok {
		return uuid.Nil, err
	}

	return uuid.Parse(userIDStr)
}
