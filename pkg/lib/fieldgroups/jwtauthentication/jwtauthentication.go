package jwtauthentication

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// JWTAuthenticationFieldGroup represents the JWTAuthenticationFieldGroup config fields
type JWTAuthenticationFieldGroup struct {
	AuthenticationType string `default:"Database"  json:"AUTHENTICATION_TYPE,omitempty" yaml:"AUTHENTICATION_TYPE,omitempty"`
	FeatureMailing     bool   `default:"false"  json:"FEATURE_MAILING" yaml:"FEATURE_MAILING"`
	JwtAuthIssuer      string `default:""  json:"JWT_AUTH_ISSUER,omitempty" yaml:"JWT_AUTH_ISSUER,omitempty"`
	JwtGetuserEndpoint string `default:""  json:"JWT_GETUSER_ENDPOINT,omitempty" yaml:"JWT_GETUSER_ENDPOINT,omitempty"`
	JwtQueryEndpoint   string `default:""  json:"JWT_QUERY_ENDPOINT,omitempty" yaml:"JWT_QUERY_ENDPOINT,omitempty"`
	JwtVerifyEndpoint  string `default:""  json:"JWT_VERIFY_ENDPOINT,omitempty" yaml:"JWT_VERIFY_ENDPOINT,omitempty"`
}

// NewJWTAuthenticationFieldGroup creates a new JWTAuthenticationFieldGroup
func NewJWTAuthenticationFieldGroup(fullConfig map[string]interface{}) (*JWTAuthenticationFieldGroup, error) {
	newJWTAuthenticationFieldGroup := &JWTAuthenticationFieldGroup{}
	defaults.Set(newJWTAuthenticationFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newJWTAuthenticationFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newJWTAuthenticationFieldGroup, nil
}
