package server

import (
	"fiber-boilerplate/database"
	"fiber-boilerplate/pkg/config"
	"fiber-boilerplate/pkg/logger"
	"fiber-boilerplate/pkg/middleware"
	"fmt"
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
