package server

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"sfcar/internal/auth_util"
)

// GRPCConfig defines the configuration that the GRPC service needs to use.
type GRPCConfig struct {
	Name          string
	Addr          string
	PublicKeyFile string
	RegisterFunc  func(server *grpc.Server)
	Logger        *zap.Logger
}

// RunGRPCServer runs a grpc server .
func RunGRPCServer(c *GRPCConfig) error {
	serverName := zap.String("name", c.Name)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("cannot listen", serverName, zap.Error(err))
	}

	// opts define the start configuration for the GRPC service.
	var opts []grpc.ServerOption
	if c.PublicKeyFile != "" {
		// Get the interceptor from auth_util package.
		in, err := auth_util.Interceptor(c.PublicKeyFile)
		if err != nil {
			c.Logger.Fatal("cannot create token interceptor in", serverName, zap.Error(err))
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}

	// Start the GRPC server with the configuration.
	s := grpc.NewServer(opts...)
	c.RegisterFunc(s)
	c.Logger.Info("server started", serverName, zap.String("addr", c.Addr))

	return s.Serve(lis)
}
