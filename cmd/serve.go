package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"smart-mess/config"
	"smart-mess/infra/conn"
	httpControllers "smart-mess/infra/http/controllers"
	httpRoutes "smart-mess/infra/http/routes"
	httpServer "smart-mess/infra/http/server"
	"smart-mess/infra/logger"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	// base context
	baseContext := context.Background()

	logger.NewLogClient(config.App().LogLevel)
	lc := logger.Client()
	dbClient := conn.Db()
	_ = dbClient

	// HttpServer
	var HttpServer = httpServer.New()

	authController := httpControllers.NewAuthController(
		baseContext,
	)

	var Routes = httpRoutes.New(
		HttpServer.Echo,
		authController,
	)

	// Spooling
	Routes.Init()
	HttpServer.Start(lc)
}
