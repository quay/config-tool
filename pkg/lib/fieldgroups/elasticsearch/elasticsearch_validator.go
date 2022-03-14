package elasticsearch

import (
	"fmt"

	"github.com/quay/config-tool/pkg/lib/shared"
)

// Validate checks the configuration settings for this field group
func (fg *ElasticSearchFieldGroup) Validate(opts shared.Options) []shared.ValidationError {

	fgName := "ElasticSearch"

	// Make empty errors
	errors := []shared.ValidationError{}

	// If not set, skip
	if fg.LogsModel != "elasticsearch" {
		return errors
	}

	// If log model config is missing
	if fg.LogsModelConfig == nil {
		newError := shared.ValidationError{
			Tags:       []string{"LOGS_MODEL_CONFIG"},
			FieldGroup: fgName,
			Message:    "LOGS_MODEL_CONFIG is required for Elasticsearch",
		}
		errors = append(errors, newError)
		return errors
	}

	// Check for elastic search config
	if fg.LogsModelConfig.ElasticsearchConfig == nil {
		newError := shared.ValidationError{
			Tags:       []string{"LOGS_MODEL_CONFIG.ELASTIC_SEARCH_CONFIG"},
			FieldGroup: fgName,
			Message:    "LOGS_MODEL_CONFIG.ELASTIC_SEARCH_CONFIG is required for Elasticsearch",
		}
		errors = append(errors, newError)
		return errors
	}

	// Check that host is available
	if fg.LogsModelConfig.ElasticsearchConfig.Host == "" {
		newError := shared.ValidationError{
			Tags:       []string{"LOGS_MODEL_CONFIG.ELASTIC_SEARCH_CONFIG.HOST"},
			FieldGroup: fgName,
			Message:    "LOGS_MODEL_CONFIG.ELASTIC_SEARCH_CONFIG.HOST is required for Elasticsearch",
		}
		errors = append(errors, newError)
	}

	// Check for port
	if fg.LogsModelConfig.ElasticsearchConfig.Port == 0 {
		newError := shared.ValidationError{
			Tags:       []string{"LOGS_MODEL_CONFIG.ELASTIC_SEARCH_CONFIG.PORT"},
			FieldGroup: fgName,
			Message:    "LOGS_MODEL_CONFIG.ELASTIC_SEARCH_CONFIG.PORT is required for Elasticsearch",
		}
		errors = append(errors, newError)
	}

	// Check for access key
	if fg.LogsModelConfig.ElasticsearchConfig.AccessKey == "" {
		newError := shared.ValidationError{
			Tags:       []string{"LOGS_MODEL_CONFIG.ELASTIC_SEARCH_CONFIG.ACCESS_KEY"},
			FieldGroup: fgName,
			Message:    "LOGS_MODEL_CONFIG.ELASTIC_SEARCH_CONFIG.ACCESS_KEY is required for Elasticsearch",
		}
		errors = append(errors, newError)
	}

	if fg.LogsModelConfig.ElasticsearchConfig.SecretKey == "" {
		newError := shared.ValidationError{
			Tags:       []string{"LOGS_MODEL_CONFIG.ELASTIC_SEARCH_CONFIG.SECRET_KEY"},
			FieldGroup: fgName,
			Message:    "LOGS_MODEL_CONFIG.ELASTIC_SEARCH_CONFIG.SECRET_KEY is required for Elasticsearch",
		}
		errors = append(errors, newError)
	}

	// Validate OAuth
	var success bool
	if opts.Mode != "testing" {
		// Get parameters to build url
		host := fg.LogsModelConfig.ElasticsearchConfig.Host
		port := fmt.Sprintf("%d", fg.LogsModelConfig.ElasticsearchConfig.Port)
		//indexPrefix := fg.LogsModelConfig.ElasticsearchConfig.IndexPrefix

		// Build url
		url := "https://" + host + ":" + port + "/" + fg.LogsModelConfig.ElasticsearchConfig.IndexPrefix + "*"
		success = shared.ValidateElasticSearchCredentials(url, fg.LogsModelConfig.ElasticsearchConfig.AccessKey, fg.LogsModelConfig.ElasticsearchConfig.SecretKey)

	} else {
		// Mock test
		success = (fg.LogsModelConfig.ElasticsearchConfig.AccessKey == "test_client_key") && (fg.LogsModelConfig.ElasticsearchConfig.SecretKey == "test_secret_key")
	}

	if !success {
		newError := shared.ValidationError{
			Tags:       []string{"LOGS_MODEL_CONFIG.ELASTIC_SEARCH_CONFIG.SECRET_KEY", "LOGS_MODEL_CONFIG.ELASTIC_SEARCH_CONFIG.SECRET_KEY"},
			FieldGroup: fgName,
			Message:    "Could not validate Elasticsearch credentials",
		}
		errors = append(errors, newError)
	}

	return errors

}
