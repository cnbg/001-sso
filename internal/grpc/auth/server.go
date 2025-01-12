package auth

import (
    "context"
    ssov1 "github.com/cnbg/001-protos/gen/go/sso"
    "google.golang.org/grpc"
)

type serverAPI struct {
    ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
    ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
    ctx context.Context,
    req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
    return &ssov1.LoginResponse{
        Token: "token123",
    }, nil
}

func (s *serverAPI) Register(
    ctx context.Context,
    req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
    panic("implement register")
}
