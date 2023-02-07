package auth

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	authpb "sfcar/auth/api/gen/v1"
)

type Service struct {
	Logger         *zap.Logger
	OpenIDResolver OpenIDResolver
	*authpb.UnimplementedAuthServiceServer
}

// OpenIDResolver resolves an authorization code
// to return the WeChat openid.
type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

// Login return the response to WeChat mini-program,the response contain
// openid,session_key.
func (s *Service) Login(ctx context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received code",
		zap.String("code", request.Code))

	openID, err := s.OpenIDResolver.Resolve(request.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "failed resolve openid: %v", err)
	}
	response := &authpb.LoginResponse{
		AccessToken: "token for openid: " + openID,
		ExpiresIn:   3600,
	}
	return response, nil
}

func (s *Service) mustEmbedUnimplementedAuthServiceServer() {

}
