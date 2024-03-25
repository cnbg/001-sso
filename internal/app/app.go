package app

import (
    "log/slog"
    grpcapp "sso/internal/app/grpc"
    "time"
)

type App struct {
    GRPCSrv *grpcapp.App
}

func New(
    log *slog.Logger,
    grpcPORT int,
    storagePath string,
    tokenTTL time.Duration,
) *App {
    // TODO: initialize storage path

    // TODO: init auth service

    grpcApp := grpcapp.New(log, grpcPORT)

    return &App{
        GRPCSrv: grpcApp,
    }
}
