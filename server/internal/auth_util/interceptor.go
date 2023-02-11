package auth_util

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"os"
	"sfcar/internal/id"
	"strings"
)

const (
	authorizationHeader = "authorization"
	bearerPrefix        = "Bearer "
)

type tokenVerifier interface {
	Verify(token string) (id.AccountID, error)
}

type interceptor struct {
	publicKey *rsa.PublicKey
	verifier  tokenVerifier
}

// Interceptor creates a grpc auth interceptor.
func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	f, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open public key file: %v", err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cannot read public key file: %v", err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("parse public key failed: %v", err)
	}

	i := &interceptor{
		publicKey: pubKey,
		verifier: &JWTTokenVerifier{
			PublicKey: pubKey,
		},
	}
	return i.HandleReq, nil
}

// HandleReq intercepts the received GRPC request and parses the token.
func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	token, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	accountID, err := i.verifier.Verify(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token invalid: %v", err)
	}
	return handler(ContextWithAccountID(ctx, accountID), req)
}

// tokenFromContext Get the token from the passed Context.
func tokenFromContext(ctx context.Context) (string, error) {
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "")
	}

	token := ""
	for _, v := range m[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) {
			token = v[len(bearerPrefix):]
		}
	}
	if token == "" {
		return "", status.Errorf(codes.Unauthenticated, "")
	}
	return token, nil
}

type accountIDKey struct{}

// ContextWithAccountID creates a context with given accountID.
func ContextWithAccountID(ctx context.Context, accountID id.AccountID) context.Context {
	return context.WithValue(ctx, accountIDKey{}, accountID)
}

// AccountIDFromContext gets account id from context.
// Returns Unauthenticated error if no account id in the context.
func AccountIDFromContext(ctx context.Context) (id.AccountID, error) {
	v := ctx.Value(accountIDKey{})
	accountID, ok := v.(id.AccountID)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "")
	}
	return accountID, nil
}
