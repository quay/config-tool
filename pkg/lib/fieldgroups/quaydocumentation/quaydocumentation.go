package quaydocumentation

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// QuayDocumentationFieldGroup represents the QuayDocumentationFieldGroup config fields
type QuayDocumentationFieldGroup struct {
	DocumentationRoot string `default:""  json:"DOCUMENTATION_ROOT,omitempty" yaml:"DOCUMENTATION_ROOT,omitempty"`
}

// NewQuayDocumentationFieldGroup creates a new QuayDocumentationFieldGroup
func NewQuayDocumentationFieldGroup(fullConfig map[string]interface{}) (*QuayDocumentationFieldGroup, error) {
	newQuayDocumentationFieldGroup := &QuayDocumentationFieldGroup{}
	defaults.Set(newQuayDocumentationFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newQuayDocumentationFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newQuayDocumentationFieldGroup, nil
}
