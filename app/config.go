package app

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type postgresConfig struct {
	LogQueries bool `toml:"log_queries"`
	URI        string
}

type miscConfig struct {
	CI bool
}

type serverConfig struct {
	AllowedHosts []string `toml:"allowed_hosts"`
	Port         string
	SecretKey    string
	ServeStatic  bool   `toml:"serve_static"`
	StaticPath   string `toml:"static_path"`
}

type smtpConfig struct {
	Address  string
	Login    string
	Password string
	Host     string
	Port     string
	UseDummy bool `toml:"use_dummy"`
}

type AppConfig struct {
	Postgres      postgresConfig
	Server        serverConfig
	SMTP          smtpConfig `toml:"smtp"`
	Miscellaneous miscConfig
}

func require(env string) string {
	envValue := os.Getenv(env)
	if envValue == "" {
		log.Fatalf("Missing environment variable %s", env)
	}
	return envValue
}

func LoadConfig(configFilePath string) (*AppConfig, error) {
	config := AppConfig{}

	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	_, err = toml.Decode(string(configFile), &config)
	if err != nil {
		return nil, err
	}

	config.SMTP.Address = os.Getenv("INKSTER_SMTP_ADDR")
	config.SMTP.Host = os.Getenv("INKSTER_SMTP_HOST")
	config.SMTP.Login = os.Getenv("INKSTER_SMTP_LOGIN")
	config.SMTP.Password = os.Getenv("INKSTER_SMTP_PASS")
	config.SMTP.Port = os.Getenv("INKSTER_SMTP_PORT")

	config.Miscellaneous.CI = os.Getenv("CI") != "" && os.Getenv("CI") == "false"

	config.Postgres.URI = require("POSTGRES_HOST")

	config.Server.Port = require("INKSTER_PORT")
	config.Server.SecretKey = require("INKSTER_SECRET_KEY")

	return &config, nil
}
