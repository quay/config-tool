package securityscanner

import (
	"github.com/creasty/defaults"
	"github.com/quay/config-tool/pkg/lib/shared"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// SecurityScannerFieldGroup represents the SecurityScannerFieldGroup config fields
type SecurityScannerFieldGroup struct {
	FeatureSecurityScanner              bool                `default:"false"  json:"FEATURE_SECURITY_SCANNER" yaml:"FEATURE_SECURITY_SCANNER"`
	SecurityScannerEndpoint             string              `default:""  json:"SECURITY_SCANNER_ENDPOINT,omitempty" yaml:"SECURITY_SCANNER_ENDPOINT,omitempty"`
	SecurityScannerIndexingInterval     *shared.IntOrString `default:"30"  json:"SECURITY_SCANNER_INDEXING_INTERVAL,omitempty" yaml:"SECURITY_SCANNER_INDEXING_INTERVAL,omitempty"`
	SecurityScannerNotifications        bool                `default:"false"  json:"FEATURE_SECURITY_NOTIFICATIONS" yaml:"FEATURE_SECURITY_NOTIFICATIONS"`
	SecurityScannerV4Endpoint           string              `default:""  json:"SECURITY_SCANNER_V4_ENDPOINT,omitempty" yaml:"SECURITY_SCANNER_V4_ENDPOINT,omitempty"`
	SecurityScannerV4NamespaceWhitelist []string            `default:"[]"  json:"SECURITY_SCANNER_V4_NAMESPACE_WHITELIST,omitempty" yaml:"SECURITY_SCANNER_V4_NAMESPACE_WHITELIST,omitempty"`
	SecurityScannerV4PSK                string              `default:"" json:"SECURITY_SCANNER_V4_PSK,omitempty" yaml:"SECURITY_SCANNER_V4_PSK,omitempty"`
}

// NewSecurityScannerFieldGroup creates a new SecurityScannerFieldGroup
func NewSecurityScannerFieldGroup(fullConfig map[string]interface{}) (*SecurityScannerFieldGroup, error) {
	newSecurityScannerFieldGroup := &SecurityScannerFieldGroup{}
	defaults.Set(newSecurityScannerFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newSecurityScannerFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newSecurityScannerFieldGroup, nil
}
