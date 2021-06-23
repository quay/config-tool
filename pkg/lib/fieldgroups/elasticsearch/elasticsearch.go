package elasticsearch

import (
	"github.com/creasty/defaults"
	"github.com/quay/config-tool/pkg/lib/shared"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// ElasticSearchFieldGroup represents the ElasticSearchFieldGroup config fields
type ElasticSearchFieldGroup struct {
	LogsModel       string                 `default:"database"  json:"LOGS_MODEL,omitempty" yaml:"LOGS_MODEL,omitempty"`
	LogsModelConfig *LogsModelConfigStruct `default:""  json:"LOGS_MODEL_CONFIG,omitempty" yaml:"LOGS_MODEL_CONFIG,omitempty"`
}

// LogsModelConfigStruct represents the LogsModelConfigStruct config fields
type LogsModelConfigStruct struct {
	KafkaConfig         *KafkaConfigStruct         `default:""  json:"kafka_config,omitempty" yaml:"kafka_config,omitempty"`
	ElasticsearchConfig *ElasticsearchConfigStruct `default:""  json:"elasticsearch_config,omitempty" yaml:"elasticsearch_config,omitempty"`
	KinesisStreamConfig *KinesisStreamConfigStruct `default:""  json:"kinesis_stream_config,omitempty" yaml:"kinesis_stream_config,omitempty"`
	Producer            string                     `default:""  json:"producer,omitempty" yaml:"producer,omitempty"`
}

// KinesisStreamConfigStruct represents the KinesisStreamConfigStruct config fields
type KinesisStreamConfigStruct struct {
	Retries            *shared.IntOrString `default:""  json:"retries,omitempty" yaml:"retries,omitempty"`
	ReadTimeout        *shared.IntOrString `default:""  json:"read_timeout,omitempty" yaml:"read_timeout,omitempty"`
	MaxPoolConnections *shared.IntOrString `default:""  json:"max_pool_connections,omitempty" yaml:"max_pool_connections,omitempty"`
	AwsRegion          string              `default:""  json:"aws_region,omitempty" yaml:"aws_region,omitempty"`
	ConnectTimeout     *shared.IntOrString `default:""  json:"connect_timeout,omitempty" yaml:"connect_timeout,omitempty"`
	AwsSecretKey       string              `default:""  json:"aws_secret_key,omitempty" yaml:"aws_secret_key,omitempty"`
	StreamName         string              `default:""  json:"stream_name,omitempty" yaml:"stream_name,omitempty"`
	AwsAccessKey       string              `default:""  json:"aws_access_key,omitempty" yaml:"aws_access_key,omitempty"`
}

// ElasticsearchConfigStruct represents the ElasticsearchConfigStruct config fields
type ElasticsearchConfigStruct struct {
	AwsRegion     string               `default:""  json:"aws_region,omitempty" yaml:"aws_region,omitempty"`
	Port          *shared.IntOrString  `default:""  json:"port,omitempty" yaml:"port,omitempty"`
	AccessKey     string               `default:""  json:"access_key,omitempty" yaml:"access_key,omitempty"`
	Host          string               `default:""  json:"host,omitempty" yaml:"host,omitempty"`
	IndexPrefix   string               `default:"logentry_"  json:"index_prefix,omitempty" yaml:"index_prefix,omitempty"`
	IndexSettings *IndexSettingsStruct `default:""  json:"index_settings,omitempty" yaml:"index_settings,omitempty"`
	UseSsl        bool                 `default:"true"  json:"use_ssl" yaml:"use_ssl"`
	SecretKey     string               `default:""  json:"secret_key,omitempty" yaml:"secret_key,omitempty"`
}

// IndexSettingsStruct represents the IndexSettings struct
type IndexSettingsStruct map[string]interface{}

// KafkaConfigStruct represents the KafkaConfigStruct config fields
type KafkaConfigStruct struct {
	Topic            string              `default:""  json:"topic,omitempty" yaml:"topic,omitempty"`
	BootstrapServers []interface{}       `default:""  json:"bootstrap_servers,omitempty" yaml:"bootstrap_servers,omitempty"`
	MaxBlockSeconds  *shared.IntOrString `default:""  json:"max_block_seconds,omitempty" yaml:"max_block_seconds,omitempty"`
}

// NewElasticSearchFieldGroup creates a new ElasticSearchFieldGroup
func NewElasticSearchFieldGroup(fullConfig map[string]interface{}) (*ElasticSearchFieldGroup, error) {
	newElasticSearchFieldGroup := &ElasticSearchFieldGroup{}
	defaults.Set(newElasticSearchFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newElasticSearchFieldGroup)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return newElasticSearchFieldGroup, nil
}
