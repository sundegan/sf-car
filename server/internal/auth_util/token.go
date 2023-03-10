package auth_util

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"sfcar/internal/id"
)

// JWTTokenVerifier verifies jwt access tokens.
type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

// Verify verifies a token and returns account id.
func (v *JWTTokenVerifier) Verify(token string) (id.AccountID, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{},
		func(*jwt.Token) (interface{}, error) {
			return v.PublicKey, nil
		})
	if err != nil {
		return "", fmt.Errorf("cannot parse token: %v", err)
	}

	if !t.Valid {
		return "", fmt.Errorf("token not valid")
	}

	claims, ok := t.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("token claim is not StandardClaims")
	}

	err = claims.Valid()
	if err != nil {
		return "", fmt.Errorf("claims not valid: %v", err)
	}
	return id.AccountID(claims.Audience), nil
}
