package config

import (
	"fmt"
	"io/ioutil"
	"log"
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

type storageBackend string

func getStorageBackend(str storageBackend) (storageBackend, error) {
	switch str {
	case "local":
		return storageBackend("local"), nil
	case "s3":
		return storageBackend("s3"), nil
	default:
		return storageBackend(""), fmt.Errorf("Unknown storage backend: %s", str)
	}
}

type storageConfig struct {
	Backend           storageBackend `toml:"backend"`
	S3AccessKey       string         `toml:"s3_access_key"`
	S3SecretAccessKey string         `toml:"-"`
	S3Bucket          string         `toml:"s3_bucket"`
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
	Storage       storageConfig  `toml:"storage"`
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

	// Fill config with environment variables
	config.Postgres.URI = envs.PgHost
	config.Server.SecretKey = envs.Secret
	if config.Storage.Backend == storageBackend("s3") {
		if envs.AWSS3SecretKey != "" {
			config.Storage.S3SecretAccessKey = envs.AWSS3SecretKey
		} else if envs.AWSSecretKey != "" {
			config.Storage.S3SecretAccessKey = envs.AWSSecretKey
		} else {
			log.Fatal("Config variable storage.backend set to s3 but no secret access key given.")
		}
	}

	// Perform validation
	if config.Storage.Backend == "" {
		config.Storage.Backend = storageBackend("local")
	} else {
		config.Storage.Backend, err = getStorageBackend(config.Storage.Backend)
		if err != nil {
			log.Printf(
				"Config variable storage.backend cannot be %s. Using local storage instead.",
				config.Storage.Backend,
			)
			config.Storage.Backend = storageBackend("local")
		}
	}

	return &config
}
