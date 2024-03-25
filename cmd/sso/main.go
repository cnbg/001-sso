package main

import (
    "log/slog"
    "os"
    "sso/internal/app"
    "sso/internal/config"
)

const (
    envLocal = "local"
    envDev   = "dev"
    envProd  = "prod"
)

func main() {
    conf := config.MustLoad()

    log := setupLogger(conf.Env)

    slog.Info("starting application", slog.Any("config", conf))

    applicaiton := app.New(log, conf.GRPC.Port, conf.StoragePath, conf.TokenTTL)

    applicaiton.GRPCSrv.MustRun()
}

func setupLogger(env string) *slog.Logger {
    var log *slog.Logger

    switch env {
    case envLocal:
        log = slog.New(
            slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
        )
    case envDev:
        log = slog.New(
            slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
        )
    case envProd:
        log = slog.New(
            slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
        )
    }

    return log
}
