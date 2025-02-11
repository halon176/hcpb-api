package main

import (
	"context"
	"hcpb-api/api"
	"hcpb-api/controllers"
)

func main() {
	ctx := context.Background()
	controllers.Migrate(ctx)
	controllers.Init(ctx)
	api.Init()

}
