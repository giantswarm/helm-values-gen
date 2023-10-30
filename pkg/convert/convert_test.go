package convert

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"sigs.k8s.io/yaml"

	"github.com/giantswarm/helm-values-gen/pkg/jsonschema"
)

func TestValidateFlags(t *testing.T) {
	testCases := []struct {
		name       string
		schemaPath string
		goldenPath string
	}{
		{
			name:       "simple",
			schemaPath: "testdata/simple.json",
			goldenPath: "testdata/simple.golden.yaml",
		},
		{
			name:       "simple with missing defaults",
			schemaPath: "testdata/simple_with_missing_defaults.json",
			goldenPath: "testdata/simple_with_missing_defaults.golden.yaml",
		},
		{
			name:       "nested objects",
			schemaPath: "testdata/nested_objects.json",
			goldenPath: "testdata/nested_objects.golden.yaml",
		},
		{
			name:       "nested arrays",
			schemaPath: "testdata/nested_arrays.json",
			goldenPath: "testdata/nested_arrays.golden.yaml",
		},
		{
			name:       "deeply nested arrays",
			schemaPath: "testdata/deeply_nested_objects.json",
			goldenPath: "testdata/deeply_nested_objects.golden.yaml",
		},
		{
			name:       "array with default",
			schemaPath: "testdata/array_with_default.json",
			goldenPath: "testdata/array_with_default.golden.yaml",
		},
		{
			name:       "referenced nested objects",
			schemaPath: "testdata/referenced_nested_objects.json",
			goldenPath: "testdata/referenced_nested_objects.golden.yaml",
		},
		{
			name:       "referenced nested objects override",
			schemaPath: "testdata/referenced_nested_objects_override.json",
			goldenPath: "testdata/referenced_nested_objects_override.golden.yaml",
		},
		{
			name:       "cluster-azure",
			schemaPath: "testdata/cluster_azure.json",
			goldenPath: "testdata/cluster_azure.golden.yaml",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := jsonschema.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("error compiling schema:\n%v", err)
			}
			defaultValues, err := BuildDefaultValuesFor(schema)
			if err != nil {
				t.Fatalf("error building default values:\n%v", err)
			}
			err = convertJsonNumbers(&defaultValues)
			if err != nil {
				t.Fatalf("error converting json numbers:\n%v", err)
			}
			golden, err := loadGolden(tc.goldenPath)
			if err != nil {
				t.Fatalf("error loading golden file:\n%v", err)
			}
			if !cmp.Equal(defaultValues, golden) {
				t.Fatalf("default values do not match golden file:\n%s", cmp.Diff(defaultValues, golden))
			}
		})
	}
}

// when comparing values parsed from yaml, numbers are parsed as float64
// but json numbers are parsed as json.Number
// this function converts all json.Number
func convertJsonNumbers(defaultValues *interface{}) error {
	defaultValuesJson, err := json.Marshal(defaultValues)
	if err != nil {
		return err
	}
	defaultValuesYaml, err := yaml.JSONToYAML(defaultValuesJson)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(defaultValuesYaml, defaultValues)
	if err != nil {
		return err
	}
	return nil
}

func loadGolden(path string) (interface{}, error) {
	yamlBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var golden interface{}
	err = yaml.Unmarshal(yamlBytes, &golden)
	if err != nil {
		return nil, err
	}

	return golden, nil
}
