package main

import (
	"context"
	"flag"
	"fmt"
	"runtime/debug"

	server "cashback_info/app/rest_server"
	"cashback_info/config"
)

var (
	port = flag.Int("port", 8040, "The server port")
)

// @title           Cashback-Info API Documentation
// @version         0.0.1
// @description     Cashback-Info API
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	flag.Parse()
	ctx := context.Background()

	postgresPool, err := config.NewPostgresPool(ctx)
	if err != nil {
		debug.PrintStack()
		panic("Couldn't connect to Postgres, error: " + err.Error())
	}

	server := server.Serve(ctx, *port, postgresPool)
	server.Run(fmt.Sprintf(":%v", *port))

	defer postgresPool.Close()
}
