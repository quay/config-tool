package apptokenauthentication

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// AppTokenAuthenticationFieldGroup represents the AppTokenAuthenticationFieldGroup config fields
type AppTokenAuthenticationFieldGroup struct {
	AuthenticationType       string `default:"Database"  json:"AUTHENTICATION_TYPE,omitempty" yaml:"AUTHENTICATION_TYPE,omitempty"`
	FeatureAppSpecificTokens bool   `default:"true"  json:"FEATURE_APP_SPECIFIC_TOKENS" yaml:"FEATURE_APP_SPECIFIC_TOKENS"`
	FeatureDirectLogin       bool   `default:"true"  json:"FEATURE_DIRECT_LOGIN" yaml:"FEATURE_DIRECT_LOGIN"`
}

// NewAppTokenAuthenticationFieldGroup creates a new AppTokenAuthenticationFieldGroup
func NewAppTokenAuthenticationFieldGroup(fullConfig map[string]interface{}) (*AppTokenAuthenticationFieldGroup, error) {
	newAppTokenAuthenticationFieldGroup := &AppTokenAuthenticationFieldGroup{}
	defaults.Set(newAppTokenAuthenticationFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newAppTokenAuthenticationFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newAppTokenAuthenticationFieldGroup, nil
}
