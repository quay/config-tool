package datamodelcache

import (
	"testing"

	"github.com/quay/config-tool/pkg/lib/shared"
)

var validateDataModelCacheTests = []struct {
	name           string
	config         map[string]interface{}
	expected       bool
	expectedErrors []shared.ValidationError
}{
	{
		"Empty",
		map[string]interface{}{},
		true,
		nil,
	},
}

func TestValidateDataModelCache(t *testing.T) {
	for _, test := range validateDataModelCacheTests {
		fg, err := NewDataModelCacheFieldGroup(test.config)
		if err != nil {
			t.Error(err)
		}

	}
}
