package actionlogarchiving

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// ActionLogArchivingFieldGroup represents the ActionLogArchivingFieldGroup config fields
type ActionLogArchivingFieldGroup struct {
	ActionLogArchiveLocation string                          `default:""  json:"ACTION_LOG_ARCHIVE_LOCATION,omitempty" yaml:"ACTION_LOG_ARCHIVE_LOCATION,omitempty"`
	ActionLogArchivePath     string                          `default:""  json:"ACTION_LOG_ARCHIVE_PATH,omitempty" yaml:"ACTION_LOG_ARCHIVE_PATH,omitempty"`
	DistributedStorageConfig *DistributedStorageConfigStruct `default:""  json:"DISTRIBUTED_STORAGE_CONFIG,omitempty" yaml:"DISTRIBUTED_STORAGE_CONFIG,omitempty"`
	FeatureActionLogRotation bool                            `default:"false"  json:"FEATURE_ACTION_LOG_ROTATION" yaml:"FEATURE_ACTION_LOG_ROTATION"`
}

// DistributedStorageConfigStruct represents the DistributedStorageConfig struct
type DistributedStorageConfigStruct map[string]interface{}

// NewActionLogArchivingFieldGroup creates a new ActionLogArchivingFieldGroup
func NewActionLogArchivingFieldGroup(fullConfig map[string]interface{}) (*ActionLogArchivingFieldGroup, error) {
	newActionLogArchivingFieldGroup := &ActionLogArchivingFieldGroup{}
	defaults.Set(newActionLogArchivingFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newActionLogArchivingFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newActionLogArchivingFieldGroup, nil
}

// NewDistributedStorageConfigStruct creates a new DistributedStorageConfigStruct
func NewDistributedStorageConfigStruct(fullConfig map[string]interface{}) (*DistributedStorageConfigStruct, error) {
	newDistributedStorageConfigStruct := DistributedStorageConfigStruct{}
	for key, value := range fullConfig {
		newDistributedStorageConfigStruct[key] = value
	}
	return &newDistributedStorageConfigStruct, nil
}
