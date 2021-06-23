package signingengine

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// SigningEngineFieldGroup represents the SigningEngineFieldGroup config fields
type SigningEngineFieldGroup struct {
	FeatureSigning         bool   `default:"false"  json:"FEATURE_SIGNING" yaml:"FEATURE_SIGNING"`
	Gpg2PrivateKeyFilename string `default:"signing-private.gpg"  json:"GPG2_PRIVATE_KEY_FILENAME,omitempty" yaml:"GPG2_PRIVATE_KEY_FILENAME,omitempty"`
	Gpg2PrivateKeyName     string `default:""  json:"GPG2_PRIVATE_KEY_NAME,omitempty" yaml:"GPG2_PRIVATE_KEY_NAME,omitempty"`
	Gpg2PublicKeyFilename  string `default:"signing-public.gpg"  json:"GPG2_PUBLIC_KEY_FILENAME,omitempty" yaml:"GPG2_PUBLIC_KEY_FILENAME,omitempty"`
	SigningEngine          string `default:""  json:"SIGNING_ENGINE,omitempty" yaml:"SIGNING_ENGINE,omitempty"`
}

// NewSigningEngineFieldGroup creates a new SigningEngineFieldGroup
func NewSigningEngineFieldGroup(fullConfig map[string]interface{}) (*SigningEngineFieldGroup, error) {
	newSigningEngineFieldGroup := &SigningEngineFieldGroup{}
	defaults.Set(newSigningEngineFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newSigningEngineFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newSigningEngineFieldGroup, nil
}
