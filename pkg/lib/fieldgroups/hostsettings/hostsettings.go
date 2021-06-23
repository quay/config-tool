package hostsettings

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// HostSettingsFieldGroup represents the HostSettingsFieldGroup config fields
type HostSettingsFieldGroup struct {
	ExternalTlsTermination bool   `default:"false"  json:"EXTERNAL_TLS_TERMINATION" yaml:"EXTERNAL_TLS_TERMINATION"`
	PreferredUrlScheme     string `default:"http"  json:"PREFERRED_URL_SCHEME,omitempty" yaml:"PREFERRED_URL_SCHEME,omitempty"`
	ServerHostname         string `default:""  json:"SERVER_HOSTNAME,omitempty" yaml:"SERVER_HOSTNAME,omitempty"`
}

// NewHostSettingsFieldGroup creates a new HostSettingsFieldGroup
func NewHostSettingsFieldGroup(fullConfig map[string]interface{}) (*HostSettingsFieldGroup, error) {
	newHostSettingsFieldGroup := &HostSettingsFieldGroup{}
	defaults.Set(newHostSettingsFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newHostSettingsFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newHostSettingsFieldGroup, nil
}
