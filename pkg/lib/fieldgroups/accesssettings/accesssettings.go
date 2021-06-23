package accesssettings

import (
	"github.com/creasty/defaults"
	"github.com/quay/config-tool/pkg/lib/shared"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// AccessSettingsFieldGroup represents the AccessSettingsFieldGroup config fields
type AccessSettingsFieldGroup struct {
	AuthenticationType             string `default:"Database"  json:"AUTHENTICATION_TYPE,omitempty" yaml:"AUTHENTICATION_TYPE,omitempty"`
	FeatureAnonymousAccess         bool   `default:"true"  json:"FEATURE_ANONYMOUS_ACCESS" yaml:"FEATURE_ANONYMOUS_ACCESS"`
	FeatureDirectLogin             bool   `default:"true"  json:"FEATURE_DIRECT_LOGIN" yaml:"FEATURE_DIRECT_LOGIN"`
	FeatureGithubLogin             bool   `default:"false"  json:"FEATURE_GITHUB_LOGIN" yaml:"FEATURE_GITHUB_LOGIN"`
	FeatureGoogleLogin             bool   `default:"false"  json:"FEATURE_GOOGLE_LOGIN" yaml:"FEATURE_GOOGLE_LOGIN"`
	HasOIDCLogin                   bool   `default:"false"  json:"-" yaml:"-"`
	FeatureInviteOnlyUserCreation  bool   `default:"false"  json:"FEATURE_INVITE_ONLY_USER_CREATION" yaml:"FEATURE_INVITE_ONLY_USER_CREATION"`
	FeaturePartialUserAutocomplete bool   `default:"true"  json:"FEATURE_PARTIAL_USER_AUTOCOMPLETE" yaml:"FEATURE_PARTIAL_USER_AUTOCOMPLETE"`
	FeatureUsernameConfirmation    bool   `default:"true"  json:"FEATURE_USERNAME_CONFIRMATION" yaml:"FEATURE_USERNAME_CONFIRMATION"`
	FeatureUserCreation            bool   `default:"true"  json:"FEATURE_USER_CREATION" yaml:"FEATURE_USER_CREATION"`
	FeatureUserLastAccessed        bool   `default:"true"  json:"FEATURE_USER_LAST_ACCESSED" yaml:"FEATURE_USER_LAST_ACCESSED"`
	FeatureUserLogAccess           bool   `default:"false"  json:"FEATURE_USER_LOG_ACCESS" yaml:"FEATURE_USER_LOG_ACCESS"`
	FeatureUserMetadata            bool   `default:"false"  json:"FEATURE_USER_METADATA" yaml:"FEATURE_USER_METADATA"`
	FeatureUserRename              bool   `default:"false"  json:"FEATURE_USER_RENAME" yaml:"FEATURE_USER_RENAME"`
	FreshLoginTimeout              string `default:"10m"  json:"FRESH_LOGIN_TIMEOUT,omitempty" yaml:"FRESH_LOGIN_TIMEOUT,omitempty"`
	UserRecoveryTokenLifetime      string `default:"30m"  json:"USER_RECOVERY_TOKEN_LIFETIME,omitempty" yaml:"USER_RECOVERY_TOKEN_LIFETIME,omitempty"`
}

// NewAccessSettingsFieldGroup creates a new AccessSettingsFieldGroup
func NewAccessSettingsFieldGroup(fullConfig map[string]interface{}) (*AccessSettingsFieldGroup, error) {
	newAccessSettingsFieldGroup := &AccessSettingsFieldGroup{}
	defaults.Set(newAccessSettingsFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newAccessSettingsFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	newAccessSettingsFieldGroup.HasOIDCLogin = shared.HasOIDCProvider(fullConfig)

	return newAccessSettingsFieldGroup, nil
}
