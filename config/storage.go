package config

import "fmt"

type StorageBackend string

const StorageAwsS3 = "s3"
const StorageLocal = "local"

func getStorageBackend(str StorageBackend) (StorageBackend, error) {
	switch str {
	case StorageLocal:
		return StorageLocal, nil
	case StorageAwsS3:
		return StorageAwsS3, nil
	default:
		return StorageBackend(""), fmt.Errorf("Unknown storage backend: %s", str)
	}
}

type storageConfig struct {
	Backend  StorageBackend `toml:"backend"`
	S3Bucket string         `toml:"s3_bucket"`
}
