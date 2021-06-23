package email

import (
	"github.com/creasty/defaults"
	"github.com/quay/config-tool/pkg/lib/shared"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// EmailFieldGroup represents the EmailFieldGroup config fields
type EmailFieldGroup struct {
	BlacklistedEmailDomains  []interface{}       `default:"[]"  yaml:"BLACKLISTED_EMAIL_DOMAINS,omitempty" json:"BLACKLISTED_EMAIL_DOMAINS,omitempty"`
	FeatureBlacklistedEmails bool                `default:"false"  yaml:"FEATURE_BLACKLISTED_EMAILS" json:"FEATURE_BLACKLISTED_EMAILS"`
	FeatureMailing           bool                `default:"false"  yaml:"FEATURE_MAILING" json:"FEATURE_MAILING"`
	MailDefaultSender        string              `default:"support@quay.io"  yaml:"MAIL_DEFAULT_SENDER,omitempty" json:"MAIL_DEFAULT_SENDER,omitempty"`
	MailPassword             string              `default:""  yaml:"MAIL_PASSWORD,omitempty" json:"MAIL_PASSWORD,omitempty"`
	MailPort                 *shared.IntOrString `default:"587"  yaml:"MAIL_PORT,omitempty" json:"MAIL_PORT,omitempty"`
	MailServer               string              `default:""  yaml:"MAIL_SERVER,omitempty" json:"MAIL_SERVER,omitempty"`
	MailUseAuth              bool                `default:"false"  yaml:"MAIL_USE_AUTH" json:"MAIL_USE_AUTH"`
	MailUsername             string              `default:""  yaml:"MAIL_USERNAME,omitempty" json:"MAIL_USERNAME,omitempty"`
	MailUseTls               bool                `default:"false"  yaml:"MAIL_USE_TLS" json:"MAIL_USE_TLS"`
	FeatureFIPS              bool                `default:"false"  yaml:"FEATURE_FIPS" json:"FEATURE_FIPS"`
}

// NewEmailFieldGroup creates a new EmailFieldGroup
func NewEmailFieldGroup(fullConfig map[string]interface{}) (*EmailFieldGroup, error) {
	newEmailFieldGroup := &EmailFieldGroup{}
	defaults.Set(newEmailFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newEmailFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newEmailFieldGroup, nil
}
