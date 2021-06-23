package ldap

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// LDAPFieldGroup represents the LDAPFieldGroup config fields
type LDAPFieldGroup struct {
	AuthenticationType        string        `default:"Database"  json:"AUTHENTICATION_TYPE,omitempty" yaml:"AUTHENTICATION_TYPE,omitempty"`
	LdapAdminDn               string        `default:""  json:"LDAP_ADMIN_DN,omitempty" yaml:"LDAP_ADMIN_DN,omitempty"`
	LdapAdminPasswd           string        `default:""  json:"LDAP_ADMIN_PASSWD,omitempty" yaml:"LDAP_ADMIN_PASSWD,omitempty"`
	LdapAllowInsecureFallback bool          `default:"false"  json:"LDAP_ALLOW_INSECURE_FALLBACK" yaml:"LDAP_ALLOW_INSECURE_FALLBACK"`
	LdapBaseDn                []interface{} `default:""  json:"LDAP_BASE_DN,omitempty" yaml:"LDAP_BASE_DN,omitempty"`
	LdapEmailAttr             string        `default:"mail"  json:"LDAP_EMAIL_ATTR,omitempty" yaml:"LDAP_EMAIL_ATTR,omitempty"`
	LdapUidAttr               string        `default:"uid"  json:"LDAP_UID_ATTR,omitempty" yaml:"LDAP_UID_ATTR,omitempty"`
	LdapUri                   string        `default:"ldap://localhost"  json:"LDAP_URI,omitempty" yaml:"LDAP_URI,omitempty"`
	LdapUserFilter            string        `default:""  json:"LDAP_USER_FILTER,omitempty" yaml:"LDAP_USER_FILTER,omitempty"`
	LdapUserRdn               []interface{} `default:"[]"  json:"LDAP_USER_RDN,omitempty" yaml:"LDAP_USER_RDN,omitempty"`
}

// NewLDAPFieldGroup creates a new LDAPFieldGroup
func NewLDAPFieldGroup(fullConfig map[string]interface{}) (*LDAPFieldGroup, error) {
	newLDAPFieldGroup := &LDAPFieldGroup{}
	defaults.Set(newLDAPFieldGroup)
	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newLDAPFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newLDAPFieldGroup, nil
}
