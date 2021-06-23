package database

import (
	"github.com/creasty/defaults"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// DatabaseFieldGroup represents the DatabaseFieldGroup config fields
type DatabaseFieldGroup struct {
	DbConnectionArgs *DbConnectionArgsStruct `default:"{}"  json:"DB_CONNECTION_ARGS,omitempty" yaml:"DB_CONNECTION_ARGS,omitempty"`
	DbUri            string                  `default:""  json:"DB_URI,omitempty" yaml:"DB_URI,omitempty"`
}

// DbConnectionArgsStruct represents the DbConnectionArgsStruct config fields
type DbConnectionArgsStruct struct {
	// MySQL arguments
	Ssl          *SslStruct `default:""  json:"ssl,omitempty" yaml:"ssl,omitempty"`
	Threadlocals bool       `default:""  json:"threadlocals,omitempty" yaml:"threadlocals,omitempty"`
	Autorollback bool       `default:""  json:"autorollback,omitempty" yaml:"autorollback,omitempty"`

	// Postgres arguments
	SslRootCert string `default:""  json:"sslrootcert,omitempty" yaml:"sslrootcert,omitempty"`
	SslMode     string `default:""  json:"sslmode,omitempty" yaml:"sslmode,omitempty"`
}

// SslStruct represents the SslStruct config fields
type SslStruct struct {
	Ca string `default:""  json:"ca,omitempty" yaml:"ca,omitempty"`
}

// NewDatabaseFieldGroup creates a new DatabaseFieldGroup
func NewDatabaseFieldGroup(fullConfig map[string]interface{}) (*DatabaseFieldGroup, error) {
	newDatabaseFieldGroup := &DatabaseFieldGroup{}
	defaults.Set(newDatabaseFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newDatabaseFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newDatabaseFieldGroup, nil
}
