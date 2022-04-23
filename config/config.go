package config

import (
	"os"
	"time"

	"github.com/joeshaw/envdecode"
)

const (
	// TemplateDir stores the name of the directory that contains templates
	TemplateDir = "templates"

	// TemplateExt stores the extension used for the template files
	TemplateExt = ".gohtml"

	// StaticDir stores the name of the directory that will serve static files
	StaticDir = "static"

	// StaticPrefix stores the URL prefix used when serving static files
	StaticPrefix = "files"
)

type environment string

const (
	// EnvLocal represents the local environment
	EnvLocal environment = "local"

	// EnvTest represents the test environment
	EnvTest environment = "test"

	// EnvDevelop represents the development environment
	EnvDevelop environment = "dev"

	// EnvProduction represents the production environment
	EnvProduction environment = "prod"
)

// SwitchEnvironment sets the environment variable used to dictate which environment the application is
// currently running in.
// This must be called prior to loading the configuration in order for it to take effect.
func SwitchEnvironment(env environment) {
	if err := os.Setenv("APP_ENVIRONMENT", string(env)); err != nil {
		panic(err)
	}
}

type (
	// Config stores complete configuration
	Config struct {
		HTTP HTTPConfig
		App  AppConfig
	}

	// HTTPConfig stores HTTP configuration
	HTTPConfig struct {
		Hostname     string        `env:"HTTP_HOSTNAME"`
		Port         uint16        `env:"HTTP_PORT,default=8000"`
		ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT,default=5s"`
		WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT,default=10s"`
		IdleTimeout  time.Duration `env:"HTTP_IDLE_TIMEOUT,default=2m"`
		TLS          struct {
			Enabled     bool   `env:"HTTP_TLS_ENABLED,default=false"`
			Certificate string `env:"HTTP_TLS_CERTIFICATE"`
			Key         string `env:"HTTP_TLS_KEY"`
		}
	}

	// AppConfig stores application configuration
	AppConfig struct {
		Name        string      `env:"APP_NAME,default=Pagoda"`
		Environment environment `env:"APP_ENVIRONMENT,default=local"`
	}
)

// GetConfig loads and returns configuration
func GetConfig() (Config, error) {
	var cfg Config
	err := envdecode.StrictDecode(&cfg)
	return cfg, err
}
