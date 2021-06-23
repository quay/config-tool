package bitbucketbuildtrigger

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// BitbucketBuildTriggerFieldGroup represents the BitbucketBuildTriggerFieldGroup config fields
type BitbucketBuildTriggerFieldGroup struct {
	BitbucketTriggerConfig *BitbucketTriggerConfigStruct `default:""  json:"BITBUCKET_TRIGGER_CONFIG,omitempty" yaml:"BITBUCKET_TRIGGER_CONFIG,omitempty"`
	FeatureBitbucketBuild  bool                          `default:"false"  json:"FEATURE_BITBUCKET_BUILD" yaml:"FEATURE_BITBUCKET_BUILD"`
	FeatureBuildSupport    bool                          `default:""  json:"FEATURE_BUILD_SUPPORT" yaml:"FEATURE_BUILD_SUPPORT"`
}

// BitbucketTriggerConfigStruct represents the BitbucketTriggerConfigStruct config fields
type BitbucketTriggerConfigStruct struct {
	ConsumerKey    string `default:""  json:"CONSUMER_KEY,omitempty" yaml:"CONSUMER_KEY,omitempty"`
	ConsumerSecret string `default:""  json:"CONSUMER_SECRET,omitempty" yaml:"CONSUMER_SECRET,omitempty"`
}

// NewBitbucketBuildTriggerFieldGroup creates a new BitbucketBuildTriggerFieldGroup
func NewBitbucketBuildTriggerFieldGroup(fullConfig map[string]interface{}) (*BitbucketBuildTriggerFieldGroup, error) {
	newBitbucketBuildTriggerFieldGroup := &BitbucketBuildTriggerFieldGroup{}
	defaults.Set(newBitbucketBuildTriggerFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newBitbucketBuildTriggerFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	return newBitbucketBuildTriggerFieldGroup, nil
}
