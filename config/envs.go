package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type envVariables struct {
	AWSSecretKey string `env:"AWS_SECRET_KEY"`
	CI           string `env:"CI"`
	Debug        string `env:"DEBUG"`
	PgHost       string `env:"PG_HOST,required"`
	Secret       string `env:"SECRET,required"`
}

func load() envVariables {
	envVars := envVariables{}

	tEnvVars := reflect.TypeOf(envVars)
	vEnvVars := reflect.ValueOf(&envVars)

	for i := 0; i < tEnvVars.NumField(); i++ {
		field := tEnvVars.Field(i)
		fieldValue := vEnvVars.Elem().Field(i)

		tag := field.Tag.Get("env")
		isRequired := false

		if tag == "" {
			fieldValue.SetString(os.Getenv(field.Name))
		} else {
			tagValues := strings.Split(tag, ",")
			if tagValues[0] != "" {
				fieldValue.SetString(os.Getenv(tagValues[0]))
			}
			if len(tagValues) > 1 && tagValues[1] == "required" {
				isRequired = true
			}
		}

		if isRequired && fieldValue.String() == "" {
			panic(fmt.Sprintf(
				"Environment variable \"%s\" required but not set",
				field.Name),
			)
		}
	}

	return envVars
}

func toBool(env string) bool {
	if env == "true" || env == "1" {
		return true
	}

	return false
}
