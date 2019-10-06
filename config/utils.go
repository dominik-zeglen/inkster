package config

import "log"

func valueOrFatal(value string, valueName string) {
	if value == "" {
		log.Fatalf("Config variable %s required but not set", valueName)
	}
}
