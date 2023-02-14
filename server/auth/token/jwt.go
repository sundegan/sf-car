package token

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// JWTTokenGen generates a JWT token.
type JWTTokenGen struct {
	PrivateKey *rsa.PrivateKey
	Issuer     string
	IssuedAt   func() time.Time
}

// GenerateToken generates a token.
func (j *JWTTokenGen) GenerateToken(accountID string, expire time.Duration) (string, error) {
	nowSec := j.IssuedAt().Unix()
	// Create the Claims.
	claims := jwt.StandardClaims{
		Issuer:    j.Issuer,
		IssuedAt:  nowSec,
		ExpiresAt: nowSec + int64(expire.Seconds()),
		Audience:  accountID,
	}
	// Create tokens without carrying a signature.
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	// Sign with a private key.
	tokenStr, err := token.SignedString(j.PrivateKey)
	if err != nil {
		log.Fatalf("cannot create jwt signature: %v", err)
		return "", err
	}
	return tokenStr, nil
}
