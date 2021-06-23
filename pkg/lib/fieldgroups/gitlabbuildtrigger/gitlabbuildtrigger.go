package gitlabbuildtrigger

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// GitLabBuildTriggerFieldGroup represents the GitLabBuildTriggerFieldGroup config fields
type GitLabBuildTriggerFieldGroup struct {
	FeatureBuildSupport bool                       `default:"false"  json:"FEATURE_BUILD_SUPPORT" yaml:"FEATURE_BUILD_SUPPORT"`
	FeatureGitlabBuild  bool                       `default:"false"  json:"FEATURE_GITLAB_BUILD" yaml:"FEATURE_GITLAB_BUILD"`
	GitlabTriggerConfig *GitlabTriggerConfigStruct `default:""  json:"GITLAB_TRIGGER_CONFIG,omitempty" yaml:"GITLAB_TRIGGER_CONFIG,omitempty"`
}

// GitlabTriggerConfigStruct represents the GitlabTriggerConfigStruct config fields
type GitlabTriggerConfigStruct struct {
	GitlabEndpoint string `default:"https://gitlab.com/"  json:"GITLAB_ENDPOINT,omitempty" yaml:"GITLAB_ENDPOINT,omitempty"`
	ClientId       string `default:""  json:"CLIENT_ID,omitempty" yaml:"CLIENT_ID,omitempty"`
	ClientSecret   string `default:""  json:"CLIENT_SECRET,omitempty" yaml:"CLIENT_SECRET,omitempty"`
}

// NewGitLabBuildTriggerFieldGroup creates a new GitLabBuildTriggerFieldGroup
func NewGitLabBuildTriggerFieldGroup(fullConfig map[string]interface{}) (*GitLabBuildTriggerFieldGroup, error) {
	newGitLabBuildTriggerFieldGroup := &GitLabBuildTriggerFieldGroup{}
	defaults.Set(newGitLabBuildTriggerFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newGitLabBuildTriggerFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newGitLabBuildTriggerFieldGroup, nil
}
