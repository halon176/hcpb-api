package main

import (
	"context"
	"hcpb-api/api"
	"hcpb-api/configs"
	"hcpb-api/controllers"
	"log/slog"
)

func main() {
	ctx := context.Background()

	shutdownTracer, err := configs.InitTracer(ctx)
	if err != nil {
		slog.Error("Failed to initialize tracer", "error", err)
	} else {
		defer func() {
			if err := shutdownTracer(ctx); err != nil {
				slog.Error("Failed to shutdown tracer", "error", err)
			}
		}()
	}

	controllers.Migrate(ctx)
	controllers.Init(ctx)
	api.Init()
}
