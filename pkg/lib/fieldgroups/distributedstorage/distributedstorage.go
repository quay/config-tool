package distributedstorage

import (
	"errors"

	"github.com/creasty/defaults"
	"github.com/quay/config-tool/pkg/lib/shared"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// DistributedStorageFieldGroup represents the DistributedStorageFieldGroup config fields
type DistributedStorageFieldGroup struct {
	DistributedStorageConfig           map[string]*DistributedStorageDefinition `default:"{}"  json:"DISTRIBUTED_STORAGE_CONFIG,omitempty" yaml:"DISTRIBUTED_STORAGE_CONFIG,omitempty"`
	DistributedStoragePreference       []string                                 `default:"[]"  json:"DISTRIBUTED_STORAGE_PREFERENCE,omitempty" yaml:"DISTRIBUTED_STORAGE_PREFERENCE,omitempty"`
	DistributedStorageDefaultLocations []string                                 `default:"[]"  json:"DISTRIBUTED_STORAGE_DEFAULT_LOCATIONS,omitempty" yaml:"DISTRIBUTED_STORAGE_DEFAULT_LOCATIONS,omitempty"`
	FeatureStorageReplication          bool                                     `default:"false"  json:"FEATURE_STORAGE_REPLICATION" yaml:"FEATURE_STORAGE_REPLICATION"`
	FeatureProxyStorage                bool                                     `default:"false"  json:"FEATURE_PROXY_STORAGE" yaml:"FEATURE_PROXY_STORAGE"`
}

// DistributedStorageDefinition represents a single storage configuration as a tuple (Name, Arguments)
type DistributedStorageDefinition struct {
	Name string                         `default:"" json:",inline" yaml:",inline"`
	Args *shared.DistributedStorageArgs `default:"" json:",inline" yaml:",inline"`
}

// NewDistributedStorageFieldGroup creates a new DistributedStorageFieldGroup
func NewDistributedStorageFieldGroup(fullConfig map[string]interface{}) (*DistributedStorageFieldGroup, error) {
	newDistributedStorageFieldGroup := &DistributedStorageFieldGroup{}
	defaults.Set(newDistributedStorageFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newDistributedStorageFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newDistributedStorageFieldGroup, nil
}

func (ds *DistributedStorageDefinition) UnmarshalYAML(value *yaml.Node) error {

	// Ensure correct shape
	if len(value.Content) != 2 || value.Content[0].Tag != "!!str" || value.Content[1].Tag != "!!map" {
		return errors.New("Incorrect format for value DISTRIBUTED_STORAGE_CONFIG")
	}

	ds.Name = value.Content[0].Value
	err := value.Content[1].Decode(&ds.Args)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}

	return nil

}

func (bm *DistributedStorageDefinition) MarshalYAML() (interface{}, error) {

	name := &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: bm.Name,
	}

	args := &yaml.Node{}
	err := args.Encode(bm.Args)
	if err != nil {
		return nil, err
	}

	node := &yaml.Node{
		Kind:    yaml.SequenceNode,
		Content: []*yaml.Node{name, args},
	}

	return node, nil

}
