package datamodelcache

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/quay/config-tool/pkg/lib/shared"
)

// Validate checks the configuration settings for this field group
func (fg *DataModelCacheFieldGroup) Validate(opts shared.Options) []shared.ValidationError {

	// Make empty errors
	errors := []shared.ValidationError{}

	if fg.Engine == "redis" {
		// Check for build log host
		if ok, err := shared.ValidateRequiredString(fg.RedisDataModelCacheConfig.Host, "HOST", "DataModelCache"); !ok {
			errors = append(errors, err)
		}

		// Build options for Redis model cache and connect
		addr := fg.RedisDataModelCacheConfig.Host
		if fg.RedisDataModelCacheConfig.Port != 0 {
			addr = addr + ":" + fmt.Sprintf("%d", fg.RedisDataModelCacheConfig.Port)
		}

		options := &redis.Options{
			Addr:     addr,
			Password: fg.RedisDataModelCacheConfig.Password,
			DB:       0,
		}
		if ok, err := shared.ValidateRedisConnection(options, "HOST", "DataModelCache"); !ok {
			errors = append(errors, err)
		}
	}

	return errors

}
