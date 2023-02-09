package auth

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	authpb "sfcar/auth/api/gen/v1"
	"sfcar/auth/dao"
)

type Service struct {
	Logger         *zap.Logger
	OpenIDResolver OpenIDResolver
	Mongo          *dao.Mongo
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

	// TODO: Using OpenID to convert to AccountID.
	// It needs to be fetched from the MongoDB database.
	accountID, err := s.Mongo.ResolveAccountID(ctx, openID)
	if err != nil {
		s.Logger.Error("failed resolve account id", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	response := &authpb.LoginResponse{
		AccessToken: "token for account id: " + accountID,
		ExpiresIn:   3600,
	}
	return response, nil
}

func (s *Service) mustEmbedUnimplementedAuthServiceServer() {

}
