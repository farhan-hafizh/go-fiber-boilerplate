package server

import (
	"fmt"
	"go-fiber-boilerplate/database"
	"go-fiber-boilerplate/pkg/config"
	"go-fiber-boilerplate/pkg/logger"
	"go-fiber-boilerplate/pkg/middleware"
	"go-fiber-boilerplate/pkg/route"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func Serve() {
	appConfig := config.AppConfig()

	logger.SetupLogger()
	logr := logger.GetLogger()

	if err := database.ConnectDB(); err != nil {
		logr.Panicf("failed database setup. error: %v", err)
	}

	app := fiber.New(config.FiberConfig())

	middleware.FiberMiddleware(app)

	db := database.GetDB()

	route.NotFoundRoute(app)
	route.GeneralRoute(app)

	group := app.Group("/api/v1")
	route.PrivateRoute(group, db)
	route.PublicRoute(group, db)

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// start shutdown goroutine
	go func() {
		// capture sigterm and other system call here
		<-sigCh
		logr.Infoln("Shutting down server...")
		_ = app.Shutdown()
	}()

	// start http server
	serverAddr := fmt.Sprintf(":%d", appConfig.Port)
	if err := app.Listen(serverAddr); err != nil {
		logr.Errorf("Oops... server is not running! error: %v", err)
	}
}
