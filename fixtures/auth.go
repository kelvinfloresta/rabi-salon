package fixtures

import (
	"rabi-salon/config"
	"rabi-salon/usecases/auth_case/role"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func SystemToken(t *testing.T) string {
	claims := jwt.MapClaims{
		"user_id": "system",
		"name":    "system",
		"email":   "system@system.com",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tk, err := token.SignedString([]byte(config.AuthSecret))
	require.Nil(t, err)

	return tk
}

func NewBackofficeToken(t *testing.T, userId string) string {
	claims := jwt.MapClaims{
		"user_id":          userId,
		"name":             "backoffice",
		"email":            "backoffice@backoffice.com",
		"role":             role.Backoffice,
		"original_user_id": userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tk, err := token.SignedString([]byte(config.AuthSecret))
	require.Nil(t, err)

	return tk
}

func BackofficeToken(t *testing.T) string {
	id := User.Create(t, nil)
	return NewBackofficeToken(t, id)
}
