package uservisiblesettings

import (
	"github.com/creasty/defaults"
	"github.com/quay/config-tool/pkg/lib/shared"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// UserVisibleSettingsFieldGroup represents the UserVisibleSettingsFieldGroup config fields
type UserVisibleSettingsFieldGroup struct {
	AvatarKind               string              `default:"local"  json:"AVATAR_KIND,omitempty" yaml:"AVATAR_KIND,omitempty"`
	Branding                 *BrandingStruct     `default:""  json:"BRANDING,omitempty" yaml:"BRANDING,omitempty"`
	ContactInfo              []interface{}       `default:"[]"  json:"CONTACT_INFO,omitempty" yaml:"CONTACT_INFO,omitempty"`
	RegistryTitle            string              `default:"Project Quay"  json:"REGISTRY_TITLE,omitempty" yaml:"REGISTRY_TITLE,omitempty"`
	RegistryTitleShort       string              `default:"Project Quay"  json:"REGISTRY_TITLE_SHORT,omitempty" yaml:"REGISTRY_TITLE_SHORT,omitempty"`
	SearchMaxResultPageCount *shared.IntOrString `default:"10"  json:"SEARCH_MAX_RESULT_PAGE_COUNT,omitempty" yaml:"SEARCH_MAX_RESULT_PAGE_COUNT,omitempty"`
	SearchResultsPerPage     *shared.IntOrString `default:"10"  json:"SEARCH_RESULTS_PER_PAGE,omitempty" yaml:"SEARCH_RESULTS_PER_PAGE,omitempty"`
	EnterpriseLogoUrl        string              `default:""  json:"ENTERPRISE_LOGO_URL,omitempty" yaml:"ENTERPRISE_LOGO_URL,omitempty"`
}

// BrandingStruct represents the BrandingStruct config fields
type BrandingStruct struct {
	Logo      string `default:"/static/img/quay-horizontal-color.svg" validate:"url" json:"logo,omitempty" yaml:"logo,omitempty"`
	FooterImg string `default:"" validate:"url" json:"footer_img,omitempty" yaml:"footer_img,omitempty"`
	FooterUrl string `default:"" validate:"url" json:"footer_url,omitempty" yaml:"footer_url,omitempty"`
}

// NewUserVisibleSettingsFieldGroup creates a new UserVisibleSettingsFieldGroup
func NewUserVisibleSettingsFieldGroup(fullConfig map[string]interface{}) (*UserVisibleSettingsFieldGroup, error) {
	newUserVisibleSettingsFieldGroup := &UserVisibleSettingsFieldGroup{}
	defaults.Set(newUserVisibleSettingsFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newUserVisibleSettingsFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newUserVisibleSettingsFieldGroup, nil
}
