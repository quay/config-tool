package timemachine

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// TimeMachineFieldGroup represents the TimeMachineFieldGroup config fields
type TimeMachineFieldGroup struct {
	DefaultTagExpiration       string        `default:"2w"  json:"DEFAULT_TAG_EXPIRATION,omitempty" yaml:"DEFAULT_TAG_EXPIRATION,omitempty"`
	FeatureChangeTagExpiration bool          `default:"true"  json:"FEATURE_CHANGE_TAG_EXPIRATION" yaml:"FEATURE_CHANGE_TAG_EXPIRATION"`
	TagExpirationOptions       []interface{} `default:"[0s, 1d, 1w, 2w, 4w]"  json:"TAG_EXPIRATION_OPTIONS,omitempty" yaml:"TAG_EXPIRATION_OPTIONS,omitempty"`
}

// NewTimeMachineFieldGroup creates a new TimeMachineFieldGroup
func NewTimeMachineFieldGroup(fullConfig map[string]interface{}) (*TimeMachineFieldGroup, error) {
	newTimeMachineFieldGroup := &TimeMachineFieldGroup{}
	defaults.Set(newTimeMachineFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newTimeMachineFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newTimeMachineFieldGroup, nil
}
