package googlelogin

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// GoogleLoginFieldGroup represents the GoogleLoginFieldGroup config fields
type GoogleLoginFieldGroup struct {
	FeatureGoogleLogin bool                     `default:"false"  json:"FEATURE_GOOGLE_LOGIN" yaml:"FEATURE_GOOGLE_LOGIN"`
	GoogleLoginConfig  *GoogleLoginConfigStruct `default:""  json:"GOOGLE_LOGIN_CONFIG,omitempty" yaml:"GOOGLE_LOGIN_CONFIG,omitempty"`
}

// GoogleLoginConfigStruct represents the GoogleLoginConfigStruct config fields
type GoogleLoginConfigStruct struct {
	ClientSecret string `default:""  json:"CLIENT_SECRET,omitempty" yaml:"CLIENT_SECRET,omitempty"`
	ClientId     string `default:""  json:"CLIENT_ID,omitempty" yaml:"CLIENT_ID,omitempty"`
}

// NewGoogleLoginFieldGroup creates a new GoogleLoginFieldGroup
func NewGoogleLoginFieldGroup(fullConfig map[string]interface{}) (*GoogleLoginFieldGroup, error) {
	newGoogleLoginFieldGroup := &GoogleLoginFieldGroup{}
	defaults.Set(newGoogleLoginFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newGoogleLoginFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newGoogleLoginFieldGroup, nil
}
