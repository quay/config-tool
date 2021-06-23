package teamsyncing

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// TeamSyncingFieldGroup represents the TeamSyncingFieldGroup config fields
type TeamSyncingFieldGroup struct {
	FeatureNonsuperuserTeamSyncingSetup bool   `default:"false"  json:"FEATURE_NONSUPERUSER_TEAM_SYNCING_SETUP" yaml:"FEATURE_NONSUPERUSER_TEAM_SYNCING_SETUP"`
	FeatureTeamSyncing                  bool   `default:"false"  json:"FEATURE_TEAM_SYNCING" yaml:"FEATURE_TEAM_SYNCING"`
	TeamResyncStaleTime                 string `default:"30m" validate:"customValidateTimePattern" json:"TEAM_RESYNC_STALE_TIME,omitempty" yaml:"TEAM_RESYNC_STALE_TIME,omitempty"`
}

// NewTeamSyncingFieldGroup creates a new TeamSyncingFieldGroup
func NewTeamSyncingFieldGroup(fullConfig map[string]interface{}) (*TeamSyncingFieldGroup, error) {
	newTeamSyncingFieldGroup := &TeamSyncingFieldGroup{}
	defaults.Set(newTeamSyncingFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newTeamSyncingFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newTeamSyncingFieldGroup, nil
}
