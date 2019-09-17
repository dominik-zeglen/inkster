package config

import (
	"io/ioutil"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/imdario/mergo"
)

const (
	// ConfigFile holds name of the config file
	ConfigFile = "config.toml"

	// OverrideConfigFile holds name of the config file used to override
	// default on
	OverrideConfigFile = "config.override.toml"
)

type debugConfig struct {
	LogQueries bool `toml:"log_queries"`
}

type postgresConfig struct {
	URI string `toml:"-"`
}

type miscConfig struct {
	CI    bool
	Debug bool
}

type serverConfig struct {
	AllowedHosts []string `toml:"allowed_hosts"`
	Port         uint16   `toml:"port"`
	SecretKey    string   `toml:"-"`
}

type awsConfig struct {
	AccessKey       string `toml:"access_key"`
	Region          string `toml:"region"`
	SecretAccessKey string `toml:"-"`
}

// Config defines application config and is propagated through every
// request handler and GraphQL query resolver
type Config struct {
	AWS           awsConfig      `toml:"aws"`
	Debug         debugConfig    `toml:"debug"`
	Miscellaneous miscConfig     `toml:"-"`
	Postgres      postgresConfig `toml:"-"`
	Server        serverConfig   `toml:"server"`
}

// Load reads the config.toml file and environment variables to create
// valid application config
func Load(configPath string) *Config {
	envs := load()

	config := Config{}

	configFile, err := ioutil.ReadFile(path.Join(configPath, ConfigFile))
	if err != nil {
		panic(err)
	}

	_, err = toml.Decode(string(configFile), &config)
	if err != nil {
		panic(err)
	}

	// Override config
	override, err := ioutil.ReadFile(path.Join(configPath, OverrideConfigFile))
	if err == nil {
		overrideConfig := Config{}
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

	config.Miscellaneous.CI = toBool(envs.CI)

	// Disable all debug options if not in debug mode
	if toBool(envs.Debug) {
		config.Debug.LogQueries = false
	}

	config.Postgres.URI = envs.PgHost
	config.Server.SecretKey = envs.Secret

	return &config
}
