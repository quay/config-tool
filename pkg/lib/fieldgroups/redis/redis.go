package redis

import (
	"github.com/creasty/defaults"
	"github.com/quay/config-tool/pkg/lib/shared"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// RedisFieldGroup represents the RedisFieldGroup config fields
type RedisFieldGroup struct {
	BuildlogsRedis  *BuildlogsRedisStruct  `default:""  json:"BUILDLOGS_REDIS,omitempty" yaml:"BUILDLOGS_REDIS,omitempty"`
	UserEventsRedis *UserEventsRedisStruct `default:""  json:"USER_EVENTS_REDIS,omitempty" yaml:"USER_EVENTS_REDIS,omitempty"`
}

// UserEventsRedisStruct represents the UserEventsRedisStruct config fields
type UserEventsRedisStruct struct {
	Password string              `default:""  json:"password,omitempty" yaml:"password,omitempty"`
	Port     *shared.IntOrString `default:""  json:"port,omitempty" yaml:"port,omitempty"`
	Host     string              `default:""  json:"host,omitempty" yaml:"host,omitempty"`
}

// BuildlogsRedisStruct represents the BuildlogsRedisStruct config fields
type BuildlogsRedisStruct struct {
	Password string              `default:""  json:"password,omitempty" yaml:"password,omitempty"`
	Port     *shared.IntOrString `default:""  json:"port,omitempty" yaml:"port,omitempty"`
	Host     string              `default:""  json:"host,omitempty" yaml:"host,omitempty"`
}

// NewRedisFieldGroup creates a new RedisFieldGroup
func NewRedisFieldGroup(fullConfig map[string]interface{}) (*RedisFieldGroup, error) {
	newRedisFieldGroup := &RedisFieldGroup{}
	defaults.Set(newRedisFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newRedisFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newRedisFieldGroup, nil
}
