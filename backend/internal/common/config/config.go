package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-yaml"
)

var (
	ErrInvalidCertPath = errors.New("invalid cert path")
	ErrInvalidKeyPath = errors.New("invalid key path")
)

type Config struct {
	App App `yaml:"app"`	
	Server Server `yaml:"server"`
	Persistence Persistence `yaml:"persistence"`
	JWT JWT `yaml:"jwt"`
	Telegram Telegram `yaml:"telegram"`
}

type App struct {
	Name string `yaml:"name" validate:"required"`
	Version string `yaml:"version" validate:"required,semver"`
	Env string `yaml:"env" validate:"required,oneof=dev local prod"`
}

type Server struct {
	HTTP struct{
		Addr string `yaml:"addr" validate:"required,hostname_port"`
		
		// There're no validation, because its optional struct.
		// If enable is true, validation of this will be in separate function.
		TLS struct{
			Enable bool `yaml:"enable" validate:"omitempty"`
			ServerCertPath string `yaml:"server_cert_path" validate:"omitempty"`
			ServerKeyPath string `yaml:"server_key_path" validate:"omitempty"`
		} `yaml:"tls"`
	} `yaml:"http"`
	Conns struct {
		ReadTimeout time.Duration `yaml:"read_timeout" validate:"required,min=100ms"`
		WriteTimeout time.Duration `yaml:"write_timeout" validate:"required,min=100ms"`
		IdleTimeout time.Duration `yaml:"idle_timeout" validate:"required,min=100ms"`
	} `yaml:"conns"`
}

type Persistence struct {
	MigrationsPath string `yaml:"migrations_path" validate:"required"`
	Postgres Postgres `yaml:"postgres" validate:"required"`
}

type Postgres struct{
	Host string `yaml:"host" validate:"required,hostname"`
	Port int `yaml:"port" validate:"required,gte=1,lte=65535"`
	SSLMode string `yaml:"sslmode" validate:"required,oneof=disable enable"`
	Auth struct{
		User string `yaml:"user" validate:"required"`
		Password string `yaml:"password" validate:"required"`
		DB string `yaml:"db" validate:"required"`
	} `yaml:"auth"`
	Conns struct{
		MaxOpens int `yaml:"max_opens" validate:"required,gte=1"`
		MaxIdles int `yaml:"max_idles" validate:"required,gte=1"`
		MaxLifetime time.Duration `yaml:"max_lifetime" validate:"required,min=1m"`
		MaxIdleTime time.Duration `yaml:"max_idle_time" validate:"required,min=1m"`
	} `yaml:"conns"`
}

type JWT struct {
	TTL time.Duration `yaml:"ttl" validate:"required,min=1m"`
	PrivateKeyPath string `yaml:"private_key_path" validate:"required"`
	PublicKeyPath string `yaml:"public_key_path" validate:"required"`
}

type Telegram struct {
	BotToken string `yaml:"bot_token" validate:"required"`
	Timeout time.Duration `yaml:"timeout" validate:"required,min=100ms"`
}

// New parse the Config from file by path
// Calling this means that env vars was loaded
func New(path string) (*Config, error) {
	bytes, err := loadBytes(path)
	if err != nil {
		return nil, fmt.Errorf("can't load config file: %w", err)
	}

	cfg, err := parseBytes(bytes)
	if err != nil {
		return nil, fmt.Errorf("can't parse confug file: %w", err)
	}

	if err := validate(cfg); err != nil {
		return nil, fmt.Errorf("config is invalid: %w", err)
	}

	return cfg, nil
}

func validate(cfg *Config) error {
	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return err
	}

	return validateTLS(cfg)
}

func validateTLS(cfg *Config) error {
	if !cfg.Server.HTTP.TLS.Enable {
		return nil
	}

	certPath := filepath.Clean(cfg.Server.HTTP.TLS.ServerCertPath)
	if stat, err := os.Stat(certPath); err != nil || stat.IsDir() {
		return fmt.Errorf("%w: '%s'", ErrInvalidCertPath, certPath)
	}

	keyPath := filepath.Clean(cfg.Server.HTTP.TLS.ServerKeyPath)
	if stat, err := os.Stat(keyPath); err != nil || stat.IsDir() {
		return fmt.Errorf("%w: '%s'", ErrInvalidKeyPath, keyPath)
	}

	return nil
}

func parseBytes(bytes []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal(bytes, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func loadBytes(path string) ([]byte, error) {
	path = filepath.Clean(path)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	content := os.ExpandEnv(string(bytes))
	return []byte(content), nil
}