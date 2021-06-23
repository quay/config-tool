package repomirror

import (
	"github.com/creasty/defaults"
	"github.com/quay/config-tool/pkg/lib/shared"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// RepoMirrorFieldGroup represents the RepoMirrorFieldGroup config fields
type RepoMirrorFieldGroup struct {
	FeatureRepoMirror        bool                `default:"false"  json:"FEATURE_REPO_MIRROR" yaml:"FEATURE_REPO_MIRROR"`
	RepoMirrorInterval       *shared.IntOrString `default:"30"  json:"REPO_MIRROR_INTERVAL,omitempty" yaml:"REPO_MIRROR_INTERVAL,omitempty"`
	RepoMirrorServerHostname string              `default:""  json:"REPO_MIRROR_SERVER_HOSTNAME,omitempty" yaml:"REPO_MIRROR_SERVER_HOSTNAME,omitempty"`
	RepoMirrorTlsVerify      bool                `default:"true"  json:"REPO_MIRROR_TLS_VERIFY" yaml:"REPO_MIRROR_TLS_VERIFY"`
}

// NewRepoMirrorFieldGroup creates a new RepoMirrorFieldGroup
func NewRepoMirrorFieldGroup(fullConfig map[string]interface{}) (*RepoMirrorFieldGroup, error) {
	newRepoMirrorFieldGroup := &RepoMirrorFieldGroup{}
	defaults.Set(newRepoMirrorFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newRepoMirrorFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newRepoMirrorFieldGroup, nil
}
