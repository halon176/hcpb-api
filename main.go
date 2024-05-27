package main

import (
	"context"
	"hcpb-api/api"
	"hcpb-api/controllers"
)

func main() {
	ctx := context.Background()
	controllers.Init(ctx)
	api.Init()

}
