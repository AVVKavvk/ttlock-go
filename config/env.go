package config

import (
	"fmt"
	"os"

	"github.com/AVVKavvk/ttlock/constants"
	logging "github.com/AVVKavvk/ttlock/utils/logger"
	"github.com/joho/godotenv"
)

var (
	log *logging.Logger
	Env *EnvConfig
)

type EnvConfig struct {
	Port string

	EnvType      string
	DocsUsername string
	DocsPassword string

	// TTLOCK

	TTLockClientID     string
	TTLockClientSecret string
}

func LoadEnvConfig() *EnvConfig {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using system environment variables")
	}

	log.Info("Loading Env Variables.........")
	Env = &EnvConfig{

		Port:    MustGetEnv("PORT"),
		EnvType: GetEnvWithDefaultValue("ENV_TYPE", "development"),

		DocsUsername: MustGetEnv("DOCS_USERNAME"),
		DocsPassword: MustGetEnv("DOCS_PASSWORD"),

		// TTLOCK
		TTLockClientID:     MustGetEnv("TTLOCK_CLIENT_ID"),
		TTLockClientSecret: MustGetEnv("TTLOCK_CLIENT_SECRET"),
	}

	log.Info("Loaded all Env variables............")
	return Env
}

// MustGetEnv fetches an env variable or panics if missing
func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " environment variable is not set")
	}
	return value
}

func GetEnvWithDefaultValue(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func init() {
	// Register this package's logger initializer
	logging.RegisterPackageLogger(func() {
		log = logging.DefaultV1Context.GetLogger(
			constants.SERVICE_NAME_FOR_LOGS+".config",
			logging.LevelDebug,
		)
	})
}
