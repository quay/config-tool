package githubbuildtrigger

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// GitHubBuildTriggerFieldGroup represents the GitHubBuildTriggerFieldGroup config fields
type GitHubBuildTriggerFieldGroup struct {
	FeatureBuildSupport bool                       `default:"false"  json:"FEATURE_BUILD_SUPPORT" yaml:"FEATURE_BUILD_SUPPORT"`
	FeatureGithubBuild  bool                       `default:"false"  json:"FEATURE_GITHUB_BUILD" yaml:"FEATURE_GITHUB_BUILD"`
	GithubTriggerConfig *GithubTriggerConfigStruct `default:""  json:"GITHUB_TRIGGER_CONFIG,omitempty" yaml:"GITHUB_TRIGGER_CONFIG,omitempty"`
}

// GithubTriggerConfigStruct represents the GithubTriggerConfigStruct config fields
type GithubTriggerConfigStruct struct {
	AllowedOrganizations []interface{} `default:"[]"  json:"ALLOWED_ORGANIZATIONS,omitempty" yaml:"ALLOWED_ORGANIZATIONS,omitempty"`
	OrgRestrict          bool          `default:"false"  json:"ORG_RESTRICT" yaml:"ORG_RESTRICT"`
	ApiEndpoint          string        `default:""  json:"API_ENDPOINT,omitempty" yaml:"API_ENDPOINT,omitempty"`
	ClientSecret         string        `default:""  json:"CLIENT_SECRET,omitempty" yaml:"CLIENT_SECRET,omitempty"`
	GithubEndpoint       string        `default:""  json:"GITHUB_ENDPOINT,omitempty" yaml:"GITHUB_ENDPOINT,omitempty"`
	ClientId             string        `default:""  json:"CLIENT_ID,omitempty" yaml:"CLIENT_ID,omitempty"`
}

// NewGitHubBuildTriggerFieldGroup creates a new GitHubBuildTriggerFieldGroup
func NewGitHubBuildTriggerFieldGroup(fullConfig map[string]interface{}) (*GitHubBuildTriggerFieldGroup, error) {
	newGitHubBuildTriggerFieldGroup := &GitHubBuildTriggerFieldGroup{}
	defaults.Set(newGitHubBuildTriggerFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newGitHubBuildTriggerFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newGitHubBuildTriggerFieldGroup, nil
}
