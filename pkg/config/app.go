package config

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	Environtment  string
	Port          int
	ReadTimeout   time.Duration
	EncryptionKey string

	// JWT Conf
	JWTSecretKey                 string
	JWTSecreteExpireMinutesCount int
}

var app = &App{}

func AppConfig() *App {
	return app
}

func LoadAppConfig() {
	app.Environtment = os.Getenv("APP_ENV")
	app.EncryptionKey = os.Getenv("APP_ENCRYPTION_SECRET")
	app.Port, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	timeOut, _ := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	app.ReadTimeout = time.Duration(timeOut) * time.Second

	app.JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	app.JWTSecreteExpireMinutesCount, _ = strconv.Atoi(os.Getenv("JWT_EXPIRE_MINUTES_COUNT"))

}

// FiberConfig func for configuration Fiber app.
func FiberConfig() fiber.Config {

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(AppConfig().ReadTimeout),
	}
}
