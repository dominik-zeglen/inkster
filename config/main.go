package config

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/imdario/mergo"
)

const (
	configFilePath         = "config.toml"
	overrideConfigFilePath = "config.override.toml"
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
	Port         uint16   `toml:"port"`
	SecretKey    string
	ServeStatic  bool   `toml:"serve_static"`
	StaticPath   string `toml:"static_path"`
}

type mailConfig struct {
	WebhookURL         string `toml:"webhook_url"`
	WebhookSecret      string
	SESAccessKey       string `toml:"ses_access_key"`
	SESSecretAccessKey string
	SESSender          string `toml:"ses_sender"`
}

type awsConfig struct {
	Region string `toml:"region"`
}

// Config defines application config and is propagated through every
// request handler and GraphQL query resolver
type Config struct {
	Postgres      postgresConfig
	Server        serverConfig
	Mail          mailConfig `toml:"mail"`
	Miscellaneous miscConfig
	AWS           awsConfig `toml:"aws"`
}

// Load reads the config.toml file and environment variables to create
// valid application config
func Load() *Config {
	envs := load()
	useProductionConfig := envs.Env == "production"

	config := Config{}

	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	_, err = toml.Decode(string(configFile), &config)
	if err != nil {
		panic(err)
	}
	if useProductionConfig {
		overrideConfig := Config{}
		override, err := ioutil.ReadFile(overrideConfigFilePath)
		if err != nil {
			panic(err)
		}

		_, err = toml.Decode(string(override), &overrideConfig)
		if err != nil {
			panic(err)
		}

		err = mergo.Merge(&overrideConfig, config)
		if err != nil {
			panic(err)
		}
		config = overrideConfig
	}

	config.Mail.WebhookSecret = envs.MailAPIKey
	config.Mail.SESSecretAccessKey = envs.MailAPIKey
	config.Miscellaneous.CI = envs.CI != "" && envs.CI != "false"
	config.Postgres.URI = envs.PgHost
	config.Server.SecretKey = envs.Secret

	if useProductionConfig {
		if config.Server.ServeStatic {
			panic(fmt.Errorf("Cannot serve static in production mode"))
		}

		if config.Mail.WebhookURL != "" && config.Mail.WebhookSecret == "" {
			panic(fmt.Errorf(("Cannot use mail webhook if no secret key was given")))
		}
	}

	return &config
}
