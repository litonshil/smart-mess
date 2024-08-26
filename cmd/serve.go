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
	"smart-mess/repositories/db"

	authservice "smart-mess/services/auth"
	txservice "smart-mess/services/transaction"
	userservice "smart-mess/services/user"
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

	dbRepo := db.NewRepository(dbClient, &lc)

	txsvc := txservice.NewDBTransaction(lc, dbRepo)
	authsvc := authservice.NewAuthService(dbRepo)
	usersvc := userservice.NewUserService(dbRepo)

	_ = txsvc

	// HttpServer
	var HttpServer = httpServer.New()

	authController := httpControllers.NewAuthController(
		baseContext,
		authsvc,
	)

	userController := httpControllers.NewUserController(
		baseContext,
		usersvc,
	)

	var Routes = httpRoutes.New(
		HttpServer.Echo,
		authController,
		userController,
	)

	// Spooling
	Routes.Init()
	HttpServer.Start(lc)
}
