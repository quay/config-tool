package githublogin

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// GitHubLoginFieldGroup represents the GitHubLoginFieldGroup config fields
type GitHubLoginFieldGroup struct {
	FeatureGithubLogin bool                     `default:"false"  json:"FEATURE_GITHUB_LOGIN" yaml:"FEATURE_GITHUB_LOGIN"`
	GithubLoginConfig  *GithubLoginConfigStruct `default:""  json:"GITHUB_LOGIN_CONFIG,omitempty" yaml:"GITHUB_LOGIN_CONFIG,omitempty"`
}

// GithubLoginConfigStruct represents the GithubLoginConfigStruct config fields
type GithubLoginConfigStruct struct {
	AllowedOrganizations []interface{} `default:"[]"  json:"ALLOWED_ORGANIZATIONS,omitempty" yaml:"ALLOWED_ORGANIZATIONS,omitempty"`
	OrgRestrict          bool          `default:"false"  json:"ORG_RESTRICT,omitempty" yaml:"ORG_RESTRICT,omitempty"`
	ApiEndpoint          string        `default:""  json:"API_ENDPOINT,omitempty" yaml:"API_ENDPOINT,omitempty"`
	GithubEndpoint       string        `default:""  json:"GITHUB_ENDPOINT,omitempty" yaml:"GITHUB_ENDPOINT,omitempty"`
	ClientId             string        `default:""  json:"CLIENT_ID,omitempty" yaml:"CLIENT_ID,omitempty"`
	ClientSecret         string        `default:""  json:"CLIENT_SECRET,omitempty" yaml:"CLIENT_SECRET,omitempty"`
}

// NewGitHubLoginFieldGroup creates a new GitHubLoginFieldGroup
func NewGitHubLoginFieldGroup(fullConfig map[string]interface{}) (*GitHubLoginFieldGroup, error) {
	newGitHubLoginFieldGroup := &GitHubLoginFieldGroup{}
	defaults.Set(newGitHubLoginFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newGitHubLoginFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	return newGitHubLoginFieldGroup, nil
}
