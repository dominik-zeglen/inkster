package config

import (
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
	Mail          mailConfig     `toml:"mail"`
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
	config.AWS.SecretAccessKey = envs.AWSSecretKey

	// Perform validation
	// --
	// Validate server
	valueOrFatal(config.Server.SecretKey, "server.secret_key")
	valueOrFatal(string(config.Server.Port), "server.port")

	// Validate storage
	if config.Storage.Backend == "" {
		config.Storage.Backend = StorageLocal
	} else {
		config.Storage.Backend, err = getStorageBackend(config.Storage.Backend)
		if err != nil {
			log.Printf(
				"Config variable storage.backend cannot be %s. Using local storage instead.",
				config.Storage.Backend,
			)
			config.Storage.Backend = StorageLocal
		}
		if config.Storage.Backend == StorageAwsS3 && envs.AWSSecretKey == "" {
			log.Fatal("Config variable storage.backend set to s3 but no aws secret access key given.")
		}
	}

	// Validate mails
	if config.Mail.Backend == "" {
		config.Mail.Backend = MailTerm
	} else {
		config.Mail.Backend, err = getMailBackend(config.Mail.Backend)
		if err != nil {
			log.Printf(
				"Config variable mail.backend cannot be %s. Using terminal output instead.",
				config.Mail.Backend,
			)
			config.Mail.Backend = MailTerm
		}
		if config.Mail.Backend == MailAwsSes && envs.AWSSecretKey == "" {
			log.Fatal("Config variable mail.backend set to ses but no aws secret access key given.")
		}
	}
	valueOrFatal(config.Mail.Sender, "mail.sender")

	return &config
}
