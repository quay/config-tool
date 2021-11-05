package datamodelcache

import (
	"errors"

	"github.com/creasty/defaults"
)

// DataModelCacheFieldGroup represents the DataModelCacheFieldGroup config fields
type DataModelCacheFieldGroup struct {
	Engine string `default:"" validate:"" json:"engine,omitempty" yaml:"engine,omitempty"`
	// FIXME(alecmerdler): Ensure YAML/JSON parsing is correct, may need to add special struct tags...
	*RedisDataModelCacheConfig
}

// RedisDataModelCacheConfig represents the RedisDataModelCacheConfig config fields
type RedisDataModelCacheConfig struct {
	Password string `default:"" validate:"" json:"password,omitempty" yaml:"password,omitempty"`
	Port     int    `default:"" validate:"" json:"port,omitempty" yaml:"port,omitempty"`
	Host     string `default:"" validate:"" json:"host,omitempty" yaml:"host,omitempty"`
}

func NewDataModelCacheFieldGroup(fullConfig map[string]interface{}) (*DataModelCacheFieldGroup, error) {
	newDataModelCacheFieldGroup := &DataModelCacheFieldGroup{}
	defaults.Set(newDataModelCacheFieldGroup)

	if value, ok := fullConfig["DATA_MODEL_CACHE_CONFIG"]; ok {
		var err error
		value := value.(map[string]interface{})
		newDataModelCacheFieldGroup.RedisDataModelCacheConfig, err = NewRedisDataModelCacheConfig(value)
		if err != nil {
			return newDataModelCacheFieldGroup, err
		}
	}

	return newDataModelCacheFieldGroup, nil
}

// NewRedisDataModelCacheConfig creates a new NewRedisDataModelCacheConfig
func NewRedisDataModelCacheConfig(cacheConfig map[string]interface{}) (*RedisDataModelCacheConfig, error) {
	newRedisDataModelCacheConfig := &RedisDataModelCacheConfig{}
	defaults.Set(newRedisDataModelCacheConfig)

	if value, ok := cacheConfig["password"]; ok {
		newRedisDataModelCacheConfig.Password, ok = value.(string)
		if !ok {
			return newRedisDataModelCacheConfig, errors.New("password must be of type string")
		}
	}

	if value, ok := cacheConfig["host"]; ok {
		newRedisDataModelCacheConfig.Host, ok = value.(string)
		if !ok {
			return newRedisDataModelCacheConfig, errors.New("host must be of type string")
		}
	}

	if value, ok := cacheConfig["port"]; ok {
		newRedisDataModelCacheConfig.Port, ok = value.(int)
		if !ok {
			return newRedisDataModelCacheConfig, errors.New("port must be of type int")
		}
	}

	return newRedisDataModelCacheConfig, nil
}
