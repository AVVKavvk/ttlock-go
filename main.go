package main

import (
	"fmt"
	"net/http"

	"github.com/AVVKavvk/ttlock/config"
	_ "github.com/AVVKavvk/ttlock/docs"
	"github.com/AVVKavvk/ttlock/server"

	logging "github.com/AVVKavvk/ttlock/utils/logger"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func HttpLoggingSkipper(ctx echo.Context) bool {
	switch ctx.Path() {
	case "/metrics":
		// Quiet the periodic metrics log
		return true
	default:

		return false
	}
}
func BasicAuth() echo.MiddlewareFunc {

	username := config.Env.DocsUsername
	password := config.Env.DocsPassword
	return middleware.BasicAuth(func(u, p string, c echo.Context) (bool, error) {
		if u == username && p == password {
			return true, nil
		}
		return false, nil
	})

}

// @title TTLock Api
// @version 1.0
// @BasePath /
func main() {
	fmt.Println("Hello TTLock")

	// Close logger
	defer logging.CloseLogger()

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using system environment variables")
	}

	config.LoadEnvConfig()

	// BasicAuth middleware
	basicAuth := BasicAuth()

	router := echo.New()

	// RequestLoggerWithConfig middleware will write the logs to gin.DefaultWriter
	router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		//  Mandatory function to handle the logged values
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			// You can format this string exactly as your original custom format
			fmt.Printf("%s [INFO] [<NONE>, %s]: {\"http_method\": \"%s\", \"http_ip\": \"%s\", \"http_status\": %d, \"http_latency\": %d, \"http_agent\": \"%s\", \"http_error\": \"%v\", \"http_path\": \"%s\", \"http_bytes_in\": %s, \"http_bytes_out\": %d}\n",
				v.StartTime.Format("20060102150405.000"), // Replaces ${time_custom}
				v.URIPath,                                // Replaces ${path}
				v.Method,                                 // Replaces ${method}
				v.RemoteIP,                               // Replaces ${remote_ip}
				v.Status,                                 // Replaces ${status}
				v.Latency.Nanoseconds(),                  // Replaces ${latency}
				v.UserAgent,                              // Replaces ${user_agent}
				v.Error,                                  // Replaces ${error}
				v.URIPath,                                // Replaces ${path}
				v.ContentLength,                          // Replaces ${bytes_in}
				v.ResponseSize,                           // Replaces ${bytes_out}
			)
			return nil
		},

		// Set necessary extraction flags to true
		LogMethod:        true,
		LogRemoteIP:      true,
		LogStatus:        true,
		LogLatency:       true,
		LogUserAgent:     true,
		LogError:         true,
		LogURIPath:       true, // For ${path}
		LogContentLength: true, // For ${bytes_in}
		LogResponseSize:  true, // For ${bytes_out}

		Skipper: HttpLoggingSkipper, //
	}))

	router.Use(middleware.Recover())

	server.AddRoutes(router)

	router.GET("/docs", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})

	router.GET("/docs/*", echoSwagger.WrapHandler, basicAuth)

	port := config.Env.Port

	router.Logger.Fatal(router.Start(":" + port))
}

func init() {
	// Initialize logger
	if err := logging.InitializeAllLoggers(); err != nil {
		panic(err)
	}
}
